package logic

import (
	"context"
	"fim_server/fim_chat/chat_models"
	"fmt"

	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatDeleteLogic {
	return &ChatDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatDeleteLogic) ChatDelete(req *types.ChatDeleteRequest) (resp *types.ChatDeleteResponse, err error) {

	var chatList []chat_models.ChatModel
	l.svcCtx.DB.Find(&chatList, req.IdList)

	var useDeleteChatList []chat_models.UserChatDeleteModel
	l.svcCtx.DB.Find(&useDeleteChatList, req.IdList)
	chatDeleteMap := map[uint]struct{}{}
	for _, model := range useDeleteChatList {
		chatDeleteMap[model.ChatID] = struct{}{}
	}

	var deleteChatList []chat_models.UserChatDeleteModel

	if len(chatList) > 0 {
		for _, model := range chatList {
			// 不是自己的聊天记录
			if !(model.SendUserID == req.UserID || model.RevUserID == req.UserID) {
				fmt.Println("不是自己的聊天记录", model.ID)
				continue
			}
			// 已经删过的聊天记录
			_, ok := chatDeleteMap[model.ID]
			if ok {
				fmt.Println("已经删除过了", model.ID)
				continue
			}
			deleteChatList = append(deleteChatList, chat_models.UserChatDeleteModel{
				UserID: req.UserID,
				ChatID: model.ID,
			})
		}
	}
	if len(deleteChatList) > 0 {
		l.svcCtx.DB.Create(&deleteChatList)
	}

	logx.Infof("已删除聊天记录 %d 条", len(deleteChatList))
	return
}
