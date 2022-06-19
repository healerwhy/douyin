package svc

import (
	"douyin/service/mq/internal/config"
	"douyin/service/rpc-user-operate/useroptservice"
	"douyin/service/rpc-video-service/videoservice"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserOptSvcRpcClient useroptservice.UserOptService
	VideoSvcRpcClient   videoservice.VideoService

	RedisCache *redis.Redis
	ScriptADD  string // 在Mqs中初始化
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserOptSvcRpcClient: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserOptServiceConf)),
		VideoSvcRpcClient:   videoservice.NewVideoService(zrpc.MustNewClient(c.VideoService)),

		RedisCache: c.RedisCacheConf.NewRedis(),
	}
}
