package svc

import (
	"fim_server/common/zrpc_interceptor"
	"fim_server/core"
	"fim_server/fim_chat/chat_rpc/chat"
	"fim_server/fim_chat/chat_rpc/types/chat_rpc"
	"fim_server/fim_group/group_rpc/groups"
	"fim_server/fim_group/group_rpc/types/group_rpc"
	"fim_server/fim_user/user_api/internal/config"
	"fim_server/fim_user/user_api/internal/middleware"
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
	ChatRpc         chat_rpc.ChatClient
	GroupRpc        group_rpc.GroupsClient
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
		ChatRpc:         chat.NewChat(zrpc.MustNewClient(c.ChatRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		GroupRpc:        groups.NewGroups(zrpc.MustNewClient(c.GroupRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		Redis:           client,
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
