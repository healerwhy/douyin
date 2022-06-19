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

	ScriptREMTag string
}

const scriptLoadREM = "local arr = redis.call('SMEMBERS', KEYS[1])  for i=1, #arr do redis.call('SREM', KEYS[1], arr[i] ) end return arr"

func NewServiceContext(c config.Config) *ServiceContext {
	ServiceContext := &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		RedisCache:  c.RedisCacheConf.NewRedis(),

		UserOptSvcRpcClient: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserOptServiceConf)),
	}

	ServiceContext.ScriptREMTag, _ = ServiceContext.RedisCache.ScriptLoad(scriptLoadREM)

	return ServiceContext
}
