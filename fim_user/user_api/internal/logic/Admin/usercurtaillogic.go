package Admin

import (
	"context"
	"errors"
	"fim_server/fim_user/user_models"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCurtailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCurtailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurtailLogic {
	return &UserCurtailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCurtailLogic) UserCurtail(req *types.UserCurtailRequest) (resp *types.UserCurtailResponse, err error) {

	var user user_models.UserModel
	err = l.svcCtx.DB.Preload("UserConfModel").Take(&user, req.UserID).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	l.svcCtx.DB.Model(&user.UserConfModel).Updates(map[string]any{
		"curtail_chat":          req.CurtailChat,
		"curtail_add_user":      req.CurtailAddUser,
		"curtail_create_group":  req.CurtailCreateGroup,
		"curtail_in_group_chat": req.CurtailInGroupChat,
	})

	return
}
