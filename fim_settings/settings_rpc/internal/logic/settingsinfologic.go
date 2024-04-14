package logic

import (
	"context"
	"encoding/json"
	"fim_server/fim_settings/settings_model"

	"fim_server/fim_settings/settings_rpc/internal/svc"
	"fim_server/fim_settings/settings_rpc/types/settings_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SettingsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSettingsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingsInfoLogic {
	return &SettingsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SettingsInfoLogic) SettingsInfo(in *settings_rpc.SettingsInfoRequest) (*settings_rpc.SettingsInfoResponse, error) {
	var settingsModel settings_model.SettingsModel
	l.svcCtx.DB.First(&settingsModel)
	byteData, _ := json.Marshal(settingsModel)
	return &settings_rpc.SettingsInfoResponse{Data: byteData}, nil
}
