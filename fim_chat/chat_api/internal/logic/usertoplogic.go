package logic

import (
	"context"
	"errors"
	"fim_server/fim_chat/chat_api/internal/svc"
	"fim_server/fim_chat/chat_api/internal/types"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTopLogic {
	return &UserTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTopLogic) UserTop(req *types.UserTopRequest) (resp *types.UserTopResponse, err error) {
	if req.UserID != req.FriendID {
		// 是否是好友
		res, err := l.svcCtx.UserRpc.IsFriend(l.ctx, &user_rpc.IsFriendRequest{
			User2: uint32(req.UserID),
			User1: uint32(req.FriendID),
		})
		if err != nil {
			return nil, err
		}
		if !res.IsFriend {
			return nil, errors.New("你们还不是好友呢")
		}
	}

	var topUser chat_models.TopUserModel
	err1 := l.svcCtx.DB.Take(&topUser, "user_id = ? and top_user_id = ?", req.UserID, req.FriendID).Error
	if err1 != nil {
		// 没有置顶
		l.svcCtx.DB.Create(&chat_models.TopUserModel{
			UserID:    req.UserID,
			TopUserID: req.FriendID,
		})
		return
	}

	// 已经有置顶了
	l.svcCtx.DB.Delete(&topUser)
	return
}
