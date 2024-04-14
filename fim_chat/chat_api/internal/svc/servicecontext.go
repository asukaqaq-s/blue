package svc

import (
	"fim_server/common/zrpc_interceptor"
	"fim_server/core"
	"fim_server/fim_chat/chat_api/internal/config"
	"fim_server/fim_chat/chat_api/internal/middleware"
	"fim_server/fim_file/file_rpc/files"
	"fim_server/fim_file/file_rpc/types/file_rpc"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/fim_user/user_rpc/users"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         user_rpc.UsersClient
	FileRpc         file_rpc.FilesClient
	Redis           *redis.Client
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		FileRpc:         files.NewFiles(zrpc.MustNewClient(c.FileRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		Redis:           client,
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
