package logic

import (
	"context"
	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListInfoLogic {
	return &UserListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListInfoLogic) UserListInfo(in *user_rpc.UserListInfoRequest) (*user_rpc.UserListInfoResponse, error) {

	fmt.Println(l.ctx.Value("clientIP"), l.ctx.Value("userID"))
	var userList []user_models.UserModel
	l.svcCtx.DB.Find(&userList, in.UserIdList)

	resp := new(user_rpc.UserListInfoResponse)
	resp.UserInfo = make(map[uint32]*user_rpc.UserInfo, 0)
	for _, i2 := range userList {
		resp.UserInfo[uint32(i2.ID)] = &user_rpc.UserInfo{
			NickName: i2.Nickname,
			Avatar:   i2.Avatar,
		}
	}
	return resp, nil
}
