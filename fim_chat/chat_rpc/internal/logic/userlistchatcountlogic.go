package logic

import (
	"context"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_chat/chat_rpc/internal/svc"
	"fim_server/fim_chat/chat_rpc/types/chat_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListChatCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListChatCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListChatCountLogic {
	return &UserListChatCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserListChatCount 查询用户发送的消息个数和接收消息的个数
func (l *UserListChatCountLogic) UserListChatCount(in *chat_rpc.UserListChatCountRequest) (*chat_rpc.UserListChatCountResponse, error) {
	type Data struct {
		UserID uint32 `gorm:"column:user_id"`
		Count  int32  `gorm:"column:count"`
	}
	var sendList []Data
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("send_user_id in ?", in.UserIdList).
		Group("send_user_id").
		Select("send_user_id as user_id", "count(id) as count").Scan(&sendList)

	var revList []Data
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("rev_user_id in ?", in.UserIdList).
		Group("rev_user_id").
		Select("rev_user_id as user_id", "count(id) as count").Scan(&revList)

	resp := new(chat_rpc.UserListChatCountResponse)
	resp.Result = map[uint32]*chat_rpc.ChatCountMessage{}

	for _, data := range sendList {
		resp.Result[data.UserID] = &chat_rpc.ChatCountMessage{
			SendMsgCount: data.Count,
		}
	}
	for _, data := range revList {
		res, ok := resp.Result[data.UserID]
		if !ok {
			resp.Result[data.UserID] = &chat_rpc.ChatCountMessage{
				RevMsgCount: data.Count,
			}
		} else {
			res.RevMsgCount = data.Count
		}
	}

	return resp, nil
}
