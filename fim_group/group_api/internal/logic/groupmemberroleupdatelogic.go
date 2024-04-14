package logic

import (
	"context"
	"errors"
	"fim_server/fim_group/group_models"

	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberRoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberRoleUpdateLogic {
	return &GroupMemberRoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberRoleUpdateLogic) GroupMemberRoleUpdate(req *types.GroupMemberRoleUpdateRequest) (resp *types.GroupMemberRoleUpdateResponse, err error) {
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.ID, req.UserID).Error
	if err != nil {
		return nil, errors.New("违规调用")
	}
	if member.Role != 1 {
		return nil, errors.New("权限错误")
	}
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.ID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("用户还不是群成员呢")
	}
	if !(req.Role == 2 || req.Role == 3) {
		return nil, errors.New("用户角色设置错误")
	}
	if member1.Role == req.Role {
		return
	}
	l.svcCtx.DB.Model(&member1).Update("role", req.Role)
	// 我记得在qq里面，群主把用户升级为管理员之后，会在群聊里面有个消息
	return
}
