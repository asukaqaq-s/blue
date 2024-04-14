package logic

import (
	"context"
	"errors"
	"fim_server/fim_user/user_models"

	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {

	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, "open_id = ?", in.OpenId).Error
	if err == nil {
		return nil, errors.New("该用户已存在")
	}
	user = user_models.UserModel{
		Nickname:       in.NickName,
		Avatar:         in.Avatar,
		Role:           int8(in.Role),
		OpenID:         in.OpenId,
		RegisterSource: in.RegisterSource,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建用户失败")
	}

	// 创建用户配置
	l.svcCtx.DB.Create(&user_models.UserConfModel{
		UserID:        user.ID,
		RecallMessage: nil,   // 撤回消息的提示内容  撤回了一条消息
		FriendOnline:  false, // 关闭好友上线提醒
		Sound:         true,  // 开启声音
		SecureLink:    false, // 关闭安全链接
		SavePwd:       false, // 不保存密码
		SearchUser:    2,     // 可以通过用户id和昵称搜索
		Verification:  2,     // 需要验证消息
		Online:        true,
	})

	return &user_rpc.UserCreateResponse{UserId: int32(user.ID)}, nil
}
