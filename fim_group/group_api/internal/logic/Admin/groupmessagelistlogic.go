package Admin

import (
	"context"
	"fim_server/common/list_query"
	"fim_server/common/models"
	"fim_server/common/models/ctype"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils"
	"time"

	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMessageListLogic {
	return &GroupMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HistoryResponse struct {
	UserID         uint          `json:"userID"`
	UserNickname   string        `json:"userNickname"`
	UserAvatar     string        `json:"userAvatar"`
	Msg            ctype.Msg     `json:"msg"`
	ID             uint          `json:"id"`
	MsgType        ctype.MsgType `json:"msgType"`
	CreatedAt      time.Time     `json:"createdAt"`
	IsMe           bool          `json:"isMe"`
	MemberNickname string        `json:"memberNickname"` // 群好友备注
}

type HistoryListResponse struct {
	List  []HistoryResponse `json:"list"`
	Count int               `json:"count"`
}

func (l *GroupMessageListLogic) GroupMessageList(req *types.GroupMessageListRequest) (resp *HistoryListResponse, err error) {

	groupMsgList, count, err := list_query.ListQuery(l.svcCtx.DB, group_models.GroupMsgModel{GroupID: req.ID}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		Preload: []string{"GroupMemberModel"},
	})

	var userIDList []uint32
	for _, model := range groupMsgList {
		userIDList = append(userIDList, uint32(model.SendUserID))
	}
	userIDList = utils.DeduplicationList(userIDList)
	userListResponse, err1 := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})

	var list = make([]HistoryResponse, 0)
	for _, model := range groupMsgList {
		info := HistoryResponse{
			UserID:    model.SendUserID,
			Msg:       model.Msg,
			ID:        model.ID,
			MsgType:   model.MsgType,
			CreatedAt: model.CreatedAt,
		}
		if model.GroupMemberModel != nil {
			info.MemberNickname = model.GroupMemberModel.MemberNickname
		}
		if err1 == nil {
			info.UserNickname = userListResponse.UserInfo[uint32(info.UserID)].NickName
			info.UserAvatar = userListResponse.UserInfo[uint32(info.UserID)].Avatar
		}
		list = append(list, info)
	}

	resp = new(HistoryListResponse)
	resp.List = list
	resp.Count = int(count)
	return
}
