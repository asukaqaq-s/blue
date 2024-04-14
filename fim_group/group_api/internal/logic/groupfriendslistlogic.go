package logic

import (
	"context"
	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupfriendsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupfriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupfriendsListLogic {
	return &GroupfriendsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupfriendsListLogic) GroupfriendsList(req *types.GroupfriendsListRequest) (resp *types.GroupfriendsListResponse, err error) {
	// 我的好友哪些在这个群里面

	// 需要去查我的好友列表
	friendResponse, err := l.svcCtx.UserRpc.FriendList(l.ctx, &user_rpc.FriendListRequest{
		User: uint32(req.UserID),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 这个群的群成员列表 组成一个map
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ?", req.ID)
	var memberMap = map[uint]bool{}
	for _, model := range memberList {
		memberMap[model.UserID] = true
	}
	resp = new(types.GroupfriendsListResponse)

	for _, info := range friendResponse.FriendList {
		resp.List = append(resp.List, types.GroupfriendsResponse{
			UserId:    uint(info.UserId),
			Avatar:    info.Avatar,
			Nickname:  info.NickName,
			IsInGroup: memberMap[uint(info.UserId)],
		})
	}
	return
}
