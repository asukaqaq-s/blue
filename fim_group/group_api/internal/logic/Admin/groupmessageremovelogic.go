package Admin

import (
	"context"
	"fim_server/fim_group/group_models"

	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMessageRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMessageRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMessageRemoveLogic {
	return &GroupMessageRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMessageRemoveLogic) GroupMessageRemove(req *types.GroupMessageRemoveRequest) (resp *types.GroupMessageRemoveResponse, err error) {
	var messageList []group_models.GroupMsgModel
	l.svcCtx.DB.Find(&messageList, "id in ?", req.IdList).Delete(&messageList)
	var userDeleteMessageList []group_models.GroupUserMsgDeleteModel
	l.svcCtx.DB.Find(&userDeleteMessageList, "msg_id in ?", req.IdList).Delete(&userDeleteMessageList)
	logx.Infof("删除聊天记录个数 %d  关联用户删除聊天记录个数 %d", len(messageList), len(userDeleteMessageList))
	return
}
