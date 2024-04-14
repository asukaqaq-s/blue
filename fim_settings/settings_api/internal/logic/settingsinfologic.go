package logic

import (
	"context"
	"fim_server/fim_settings/settings_model"

	"fim_server/fim_settings/settings_api/internal/svc"
	"fim_server/fim_settings/settings_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SettingsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettingsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingsInfoLogic {
	return &SettingsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SettingsInfoLogic) SettingsInfo(req *types.SettingsInfoRequest) (resp settings_model.SettingsModel, err error) {

	// 有且只有一条记录
	// 在查询的时候，查不到就添加一条记录
	// 在系统启动的时候，查一下有没有，没有就加一条
	l.svcCtx.DB.First(&resp)
	resp.QQ.Key = "******"
	resp.QQ.WebPath = resp.QQ.GetPath()

	return
}
