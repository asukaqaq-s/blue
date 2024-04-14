package Admin

import (
	"context"
	"errors"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatAdminSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatAdminSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatAdminSessionLogic {
	return &ChatAdminSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ChatAdminSession 作为接收者的用户，哪些人和他聊过
func (l *ChatAdminSessionLogic) ChatAdminSession(req *types.ChatAdminSessionRequest) (resp *types.ChatAdminSessionResponse, err error) {

	var sendUserIDList []uint32
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("rev_user_id = ?", req.RevUserID).
		Group("send_user_id").
		Select("send_user_id").Scan(&sendUserIDList)

	userList, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: sendUserIDList,
	})
	if err != nil {
		return nil, errors.New("用户服务错误")
	}
	resp = new(types.ChatAdminSessionResponse)
	for u, info := range userList.UserInfo {
		resp.List = append(resp.List, types.UserInfo{
			UserID:   uint(u),
			Avatart:  info.Avatar,
			Nickname: info.NickName,
		})
	}
	resp.Count = len(resp.List)
	return
}
