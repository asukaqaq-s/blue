package svc

import (
	"fim_server/common/log_stash"
	"fim_server/common/zrpc_interceptor"
	"fim_server/core"
	"fim_server/fim_logs/logs_api/internal/config"
	"fim_server/fim_logs/logs_api/internal/middleware"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/fim_user/user_rpc/users"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         user_rpc.UsersClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
	KqPusherClient  *kq.Pusher
	ActionLogs      *log_stash.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)

	kqClient := kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
		KqPusherClient:  kqClient,
		ActionLogs:      log_stash.NewActionPusher(kqClient, c.Name),
	}
}
