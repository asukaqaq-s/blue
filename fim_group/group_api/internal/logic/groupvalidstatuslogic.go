package logic

import (
	"context"
	"errors"
	"fim_server/fim_group/group_models"

	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidStatusLogic {
	return &GroupValidStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidStatusLogic) GroupValidStatus(req *types.GroupValidStatusRequest) (resp *types.GroupValidStatusResponse, err error) {

	var groupValidModel group_models.GroupVerifyModel
	err = l.svcCtx.DB.Take(&groupValidModel, req.ValidID).Error
	if err != nil {
		return nil, errors.New("不存在的验证记录")
	}
	switch req.Status {
	case 1, 2, 3:
		if groupValidModel.Status != 0 {
			return nil, errors.New("已经处理过该验证请求了")
		}
	case 4:
		if groupValidModel.Status == 0 {
			return nil, errors.New("只能删除处理过的请求")
		}
	default:
		return nil, errors.New("错误的状态")
	}

	// 判断我有没有权限处理这个验证请求
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.UserID, groupValidModel.GroupID).Error
	if err != nil {
		return nil, errors.New("没有处理该操作的权限")
	}
	if !(member.Role == 1 || member.Role == 2) {
		return nil, errors.New("没有处理该操作的权限")
	}

	switch req.Status {
	case 0: // 未操作
		return
	case 1: // 同意
		// 将用户加到群里面去
		var member1 = group_models.GroupMemberModel{
			GroupID: groupValidModel.GroupID,
			UserID:  groupValidModel.UserID,
			Role:    3,
		}
		l.svcCtx.DB.Create(&member1)
	case 2: // 拒绝
	case 3: // 忽略

	case 4: // 删掉
		l.svcCtx.DB.Delete(&groupValidModel)
		return

	}

	l.svcCtx.DB.Model(&groupValidModel).UpdateColumn("status", req.Status)

	return
}
