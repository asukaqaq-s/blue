package svc

import (
	"fim_server/common/zrpc_interceptor"
	"fim_server/core"
	"fim_server/fim_file/file_api/internal/config"
	"fim_server/fim_file/file_api/internal/middleware"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/fim_user/user_rpc/users"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	UserRpc         user_rpc.UsersClient
	DB              *gorm.DB
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		DB:              mysqlDb,
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
