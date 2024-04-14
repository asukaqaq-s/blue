package svc

import (
	"fim_server/common/log_stash"
	"fim_server/common/zrpc_interceptor"
	"fim_server/core"
	"fim_server/fim_auth/auth_api/internal/config"
	"fim_server/fim_settings/settings_rpc/settings"
	"fim_server/fim_settings/settings_rpc/types/settings_rpc"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/fim_user/user_rpc/users"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	Redis          *redis.Client
	UserRpc        user_rpc.UsersClient
	SettingsRpc    settings_rpc.SettingsClient
	KqPusherClient *kq.Pusher
	ActionLogs     *log_stash.Pusher
	RuntimeLogs    *log_stash.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	kqClient := kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
	return &ServiceContext{
		Config:         c,
		DB:             mysqlDb,
		Redis:          client,
		UserRpc:        users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		SettingsRpc:    settings.NewSettings(zrpc.MustNewClient(c.SettingsRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		KqPusherClient: kqClient,
		ActionLogs:     log_stash.NewActionPusher(kqClient, c.Name),
		RuntimeLogs:    log_stash.NewRuntimePusher(kqClient, c.Name),
	}
}
