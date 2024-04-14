package svc

import (
	"fim_server/core"
	"fim_server/fim_settings/settings_api/internal/config"
	"fim_server/fim_settings/settings_api/internal/middleware"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
