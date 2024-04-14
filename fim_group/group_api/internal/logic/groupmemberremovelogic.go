package logic

import (
	"context"
	"errors"
	"fim_server/fim_group/group_models"

	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMemberRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMemberRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMemberRemoveLogic {
	return &GroupMemberRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMemberRemoveLogic) GroupMemberRemove(req *types.GroupMemberRemoveRequest) (resp *types.GroupMemberRemoveResponse, err error) {
	// 谁能调这个接口 必须得是这个群的成员
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.ID, req.UserID).Error
	if err != nil {
		return nil, errors.New("违规调用")
	}
	// 用户自己退群
	if req.UserID == req.MemberID {
		// 自己不能是群主 群主不能退群，群主只能解散群
		if member.Role == 1 {
			return nil, errors.New("群主不能退群，只能解散群聊")
		}
		// 把member中的与这个用户的记录删掉就好了
		l.svcCtx.DB.Delete(&member)
		// 给群验证表里面加条记录
		l.svcCtx.DB.Create(&group_models.GroupVerifyModel{
			GroupID: member.GroupID,
			UserID:  req.UserID,
			Type:    2, // 退群
		})
		return

	}
	// 把用户踢出群聊
	if !(member.Role == 1 || member.Role == 2) {
		return nil, errors.New("违规调用")
	}
	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.ID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("该用户不是群成员")
	}
	// 群主可以踢管理员和用户
	// 管理员只能踢用户
	if !(member.Role == 1 && (member1.Role == 2 || member1.Role == 3) || member.Role == 2 && member1.Role == 3) {
		return nil, errors.New("角色错误")
	}
	err = l.svcCtx.DB.Delete(&member1).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("群成员移出失败")
	}
	return
}
