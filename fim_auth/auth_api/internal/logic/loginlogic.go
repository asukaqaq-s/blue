package logic

import (
	"context"
	"errors"
	"fim_server/fim_auth/auth_api/internal/svc"
	"fim_server/fim_auth/auth_api/internal/types"
	"fim_server/fim_auth/auth_models"
	"fim_server/utils/jwts"
	"fim_server/utils/pwd"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "id = ? ", req.UserName).Error
	l.svcCtx.ActionLogs.IsRequest()
	//l.svcCtx.ActionLogs.IsResponse()
	l.svcCtx.ActionLogs.IsHeaders()
	l.svcCtx.ActionLogs.Info("用户登录操作")
	l.svcCtx.ActionLogs.SetItem("nickName", req.UserName)
	defer l.svcCtx.ActionLogs.Save(l.ctx)
	if err != nil {
		l.svcCtx.ActionLogs.Err(req.UserName + " 用户名不存在")
		err = errors.New("用户名或密码错误")
		return
	}

	if !pwd.CheckPwd(user.Pwd, req.Password) {
		l.svcCtx.ActionLogs.SetItem("password", req.Password)
		l.svcCtx.ActionLogs.Err(req.UserName + " 密码错误")
		err = errors.New("用户名或密码错误")
		return
	}
	// 判断用户的注册来源，第三方登录来的不能通过用户名密码登录

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err)
		l.svcCtx.ActionLogs.SetItem("error", err.Error())
		l.svcCtx.ActionLogs.Err("服务内部错误")
		l.svcCtx.RuntimeLogs.SetItemErr("xxx", err.Error())
		err = errors.New("服务内部错误")
		return
	}
	ctx := context.WithValue(l.ctx, "userID", fmt.Sprintf("%d", user.ID))
	l.svcCtx.ActionLogs.Info("用户登录成功")
	l.svcCtx.ActionLogs.SetCtx(ctx)

	return &types.LoginResponse{Token: token}, nil
}
