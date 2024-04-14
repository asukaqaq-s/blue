package logic

import (
	"context"
	"errors"
	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"
	"fim_server/fim_group/group_models"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupProhibitionUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupProhibitionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupProhibitionUpdateLogic {
	return &GroupProhibitionUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GroupProhibitionUpdate 设置用户的禁言
func (l *GroupProhibitionUpdateLogic) GroupProhibitionUpdate(req *types.GroupProhibitionUpdateRequest) (resp *types.GroupProhibitionUpdateResponse, err error) {
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, req.UserID).Error
	if err != nil {
		return nil, errors.New("当前用户错误")
	}
	if !(member.Role == 1 || member.Role == 2) {
		return nil, errors.New("当前用户角色错误")
	}

	var member1 group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member1, "group_id = ? and user_id = ?", req.GroupID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("目标用户不是群成员")
	}
	if !((member.Role == 1 && member1.Role == 2 || member1.Role == 3) || (member.Role == 2 && member1.Role == 3)) {
		return nil, errors.New("角色错误")
	}

	l.svcCtx.DB.Model(&member1).Update("prohibition_time", req.ProhibitionTime)

	// 利用redis的过期时间去做这个禁言时间
	key := fmt.Sprintf("prohibition__%d", member1.ID)
	if req.ProhibitionTime != nil {
		// 给redis设置一个key，过期时间是xxxx
		l.svcCtx.Redis.Set(key, "1", time.Duration(*req.ProhibitionTime)*time.Minute)
	} else {
		l.svcCtx.Redis.Del(key)
	}
	return
}
