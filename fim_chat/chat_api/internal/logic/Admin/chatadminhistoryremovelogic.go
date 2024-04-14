package Admin

import (
	"context"
	"fim_server/fim_chat/chat_models"

	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatAdminHistoryRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatAdminHistoryRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatAdminHistoryRemoveLogic {
	return &ChatAdminHistoryRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatAdminHistoryRemoveLogic) ChatAdminHistoryRemove(req *types.ChatAdminHistoryRemoveRequest) (resp *types.ChatAdminHistoryRemoveResponse, err error) {
	var msgList []chat_models.ChatModel
	l.svcCtx.DB.Find(&msgList, "id in ?", req.IdList).Delete(&msgList)
	logx.Infof("删除聊天记录个数 %d", len(msgList))
	var userChatDeleteList []chat_models.UserChatDeleteModel
	l.svcCtx.DB.Find(&userChatDeleteList, "chat_id in ?", req.IdList).Delete(&userChatDeleteList)
	logx.Infof("删除关联用户删除聊天记录个数 %d", len(userChatDeleteList))
	return
}
