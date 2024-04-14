package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fim_server/common/models/ctype"
	"fim_server/fim_auth/auth_models"
	"fim_server/fim_settings/settings_rpc/types/settings_rpc"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils/jwts"
	"fim_server/utils/open_login"
	"fmt"

	"fim_server/fim_auth/auth_api/internal/svc"
	"fim_server/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {

	type OpenInfo struct {
		Nickname string
		OpenID   string
		Avatar   string
	}

	var info OpenInfo

	// 去调系统配置的rpc方法
	settingsInfoRes, err := l.svcCtx.SettingsRpc.SettingsInfo(context.Background(), &settings_rpc.SettingsInfoRequest{})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("系统配置服务异常")
	}
	type SettingsInfo struct {
		QQ ctype.QQType `json:"qq"`
	}
	var settingsInfo SettingsInfo
	err = json.Unmarshal(settingsInfoRes.Data, &settingsInfo)
	if err != nil {
		logx.Error(err)
		logx.Error(string(settingsInfoRes.Data))
		return nil, errors.New("系统配置响应异常")
	}
	fmt.Println(settingsInfo)

	switch req.Flag {
	case "qq":
		qqInfo, openError := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    settingsInfo.QQ.AppID,
			AppKey:   settingsInfo.QQ.Key,
			Redirect: settingsInfo.QQ.Redirect,
		})
		info = OpenInfo{
			OpenID:   qqInfo.OpenID,
			Nickname: qqInfo.Nickname,
			Avatar:   qqInfo.Avatar,
		}
		err = openError
	default:
		err = errors.New("不支持的三方登录")
	}
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		// 注册逻辑
		fmt.Println("注册服务")
		res, err := l.svcCtx.UserRpc.UserCreate(l.ctx, &user_rpc.UserCreateRequest{
			NickName:       info.Nickname,
			Password:       "",
			Role:           2,
			Avatar:         info.Avatar,
			OpenId:         info.OpenID,
			RegisterSource: "qq",
		})
		if err != nil {
			logx.Error(err)
			return nil, errors.New("登录失败")
		}
		user.Model.ID = uint(res.UserId)
		user.Role = 2
		user.Nickname = info.Nickname
	}

	// 登录逻辑
	token, err1 := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err1 != nil {
		logx.Error(err1)
		err = errors.New("服务内部错误")
		return nil, err1
	}
	return &types.LoginResponse{Token: token}, nil

	return
}
