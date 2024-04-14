package logic

import (
	"context"
	"fim_server/fim_group/group_models"

	"fim_server/fim_group/group_rpc/internal/svc"
	"fim_server/fim_group/group_rpc/types/group_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGroupSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGroupSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupSearchLogic {
	return &UserGroupSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGroupSearchLogic) UserGroupSearch(in *group_rpc.UserGroupSearchRequest) (resp *group_rpc.UserGroupSearchResponse, err error) {
	type Data struct {
		UserID uint32 `gorm:"column:user_id"`
		Count  uint32 `gorm:"column:count"`
	}
	var data []Data
	switch in.Mode {
	case 1: // 查用户创建的个数
		l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
			Where("user_id in ? and role = ?", in.UserIdList, 1).
			Group("user_id").
			Select("user_id", "count(id) as count").Scan(&data)
	case 2: // 查用户加人群聊的个数
		l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
			Where("user_id in ?", in.UserIdList).
			Group("user_id").
			Select("user_id", "count(id) as count").Scan(&data)
	}
	var groupUserMap = map[uint32]uint32{}
	for _, u2 := range data {
		groupUserMap[u2.UserID] = u2.Count
	}
	resp = new(group_rpc.UserGroupSearchResponse)
	resp.Result = map[uint32]int32{}
	for _, uid := range in.UserIdList {
		resp.Result[uid] = int32(groupUserMap[uid])
	}
	return resp, nil
}
