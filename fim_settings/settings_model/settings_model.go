package settings_model

import (
	"fim_server/common/models"
	"fim_server/common/models/ctype"
)

type SettingsModel struct {
	models.Model
	Site ctype.SiteType `json:"site"`
	QQ   ctype.QQType   `json:"qq"`
}
