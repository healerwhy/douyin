package svc

import (
	"douyin/service/asynqJob/server/internal/config"
	"douyin/service/rpc-user-operate/useroptservice"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	RedisCache  *redis.Redis

	UserOptSvcRpcClient useroptservice.UserOptService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		RedisCache:  c.RedisCacheConf.NewRedis(),

		UserOptSvcRpcClient: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserOptServiceConf)),
	}
}
