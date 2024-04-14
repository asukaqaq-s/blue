package logic

import (
	"context"
	"errors"
	"fim_server/fim_user/user_models"

	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBaseInfoLogic {
	return &UserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserBaseInfoLogic) UserBaseInfo(in *user_rpc.UserBaseInfoRequest) (*user_rpc.UserBaseInfoResponse, error) {
	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, in.UserId).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &user_rpc.UserBaseInfoResponse{
		UserId:   in.UserId,
		NickName: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}
