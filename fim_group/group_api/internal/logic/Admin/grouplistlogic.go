package Admin

import (
	"context"
	"fim_server/common/list_query"
	"fim_server/common/models"
	"fim_server/common/models/ctype"
	"fim_server/fim_group/group_api/internal/svc"
	"fim_server/fim_group/group_api/internal/types"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupListRequest) (resp *types.GroupListResponse, err error) {
	list, count, err := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Key:   req.Key,
			Sort:  "created_at desc",
		},
		Likes:   []string{"title"},
		Preload: []string{"MemberList", "GroupMsgList"},
	})
	var userIDList []uint32
	for _, model := range list {
		for _, memberModel := range model.MemberList {
			userIDList = append(userIDList, uint32(memberModel.UserID))
		}
	}
	userIDList = utils.DeduplicationList(userIDList)

	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	var userInfoMap = map[uint]ctype.UserInfo{}
	if err == nil {
		for u, info := range userListResponse.UserInfo {
			userInfoMap[uint(u)] = ctype.UserInfo{
				ID:       uint(u),
				NickName: info.NickName,
				Avatar:   info.Avatar,
			}
		}
	} else {
		logx.Error(err)
	}

	var userOnlineMap = map[uint]bool{}
	userOnline, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	if err == nil {
		for _, u := range userOnline.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logx.Error(err)
	}

	resp = new(types.GroupListResponse)
	for _, model := range list {
		info := types.GroupListInfoResponse{
			ID:           model.ID,
			CreatedAt:    model.CreatedAt.String(),
			Title:        model.Title,
			Abstract:     model.Abstract,
			Avatar:       model.Avatar,
			MemberCount:  len(model.MemberList),
			MessageCount: len(model.GroupMsgList),
			Creater: types.UserInfo{
				UserID:   model.Creator,
				Avatart:  userInfoMap[model.Creator].Avatar,
				Nickname: userInfoMap[model.Creator].NickName,
			},
		}

		var adminList []types.UserInfo

		for _, memberModel := range model.MemberList {
			_, ok := userOnlineMap[memberModel.UserID]
			if ok {
				info.MemberOnlineCount++
			}
			if memberModel.Role == 2 {
				adminList = append(adminList, types.UserInfo{
					UserID:   memberModel.UserID,
					Avatart:  userInfoMap[memberModel.UserID].Avatar,
					Nickname: userInfoMap[memberModel.UserID].NickName,
				})
			}
		}
		info.AdminList = adminList

		resp.List = append(resp.List, info)
	}
	resp.Count = int(count)
	return
}
