package Admin

import (
	"context"
	"errors"
	"fim_server/common/list_query"
	"fim_server/common/models"
	"fim_server/common/models/ctype"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils"

	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatAdminHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatAdminHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatAdminHistoryLogic {
	return &ChatAdminHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChatHistory struct {
	ID        uint             `json:"id"`
	SendUser  ctype.UserInfo   `json:"sendUser"`
	RevUser   ctype.UserInfo   `json:"revUser"`
	IsMe      bool             `json:"isMe"`       // 哪条消息是我发的
	CreatedAt string           `json:"created_at"` // 消息时间
	Msg       ctype.Msg        `json:"msg"`
	SystemMsg *ctype.SystemMsg `json:"systemMsg"`
}

type ChatHistoryResponse struct {
	List  []ChatHistory `json:"list"`
	Count int64         `json:"count"`
}

// ChatAdminHistory 用户与用户的聊天记录
func (l *ChatAdminHistoryLogic) ChatAdminHistory(req *types.ChatAdminHistoryRequest) (resp *ChatHistoryResponse, err error) {
	chatList, count, _ := list_query.ListQuery(l.svcCtx.DB, chat_models.ChatModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		//Debug: true,
		Where: l.svcCtx.DB.Where("((send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?))",
			req.SendUserID, req.RevUserID, req.RevUserID, req.SendUserID),
	})

	var userIDList []uint32
	for _, model := range chatList {
		userIDList = append(userIDList, uint32(model.SendUserID))
		userIDList = append(userIDList, uint32(model.RevUserID))
	}

	// 去重
	userIDList = utils.DeduplicationList(userIDList)
	// 去调用户服务的rpc方法，获取用户信息 {用户id：{用户信息}}

	response, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}

	var list = make([]ChatHistory, 0)
	for _, model := range chatList {

		sendUser := ctype.UserInfo{
			ID:       model.SendUserID,
			NickName: response.UserInfo[uint32(model.SendUserID)].NickName,
			Avatar:   response.UserInfo[uint32(model.SendUserID)].Avatar,
		}
		revUser := ctype.UserInfo{
			ID:       model.RevUserID,
			NickName: response.UserInfo[uint32(model.RevUserID)].NickName,
			Avatar:   response.UserInfo[uint32(model.RevUserID)].Avatar,
		}

		info := ChatHistory{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.String(),
			SendUser:  sendUser,
			RevUser:   revUser,
			Msg:       model.Msg,
			SystemMsg: model.SystemMsg,
		}

		if info.SendUser.ID == req.RevUserID {
			info.IsMe = true
		}

		list = append(list, info)
	}

	resp = &ChatHistoryResponse{
		List:  list,
		Count: count,
	}
	return
}
