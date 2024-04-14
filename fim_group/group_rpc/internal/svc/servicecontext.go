package svc

import (
	"fim_server/core"
	"fim_server/fim_group/group_rpc/internal/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.RedisConf.Addr, c.RedisConf.Pwd, c.RedisConf.DB)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
		Redis:  client,
	}
}
