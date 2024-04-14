package Admin

import (
	"context"
	"errors"
	"fim_server/common/models/ctype"
	"fim_server/fim_settings/settings_model"

	"fim_server/fim_settings/settings_api/internal/svc"
	"fim_server/fim_settings/settings_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SettingsInfoUpdadeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettingsInfoUpdadeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingsInfoUpdadeLogic {
	return &SettingsInfoUpdadeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type SettingsInfoUpdateRequest struct {
	Site ctype.SiteType `json:"site"`
	QQ   ctype.QQType   `json:"qq"`
}

func (l *SettingsInfoUpdadeLogic) SettingsInfoUpdade(req *SettingsInfoUpdateRequest) (resp *types.SettingsInfoUpdateResponse, err error) {

	var settingsModel settings_model.SettingsModel
	l.svcCtx.DB.First(&settingsModel)

	if req.QQ.Key == "******" {
		req.QQ.Key = settingsModel.QQ.Key
	}
	err = l.svcCtx.DB.Model(&settingsModel).Updates(settings_model.SettingsModel{
		Site: req.Site,
		QQ:   req.QQ,
	}).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("修改失败")
	}

	return
}
