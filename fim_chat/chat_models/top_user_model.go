package chat_models

import "fim_server/common/models"

// TopUserModel 置顶用户表
type TopUserModel struct {
	models.Model
	UserID    uint `json:"userID"`
	TopUserID uint `json:"topUserID"`
}
