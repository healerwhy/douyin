package svc

import (
	"douyin/service/mq/internal/config"
	"douyin/service/rpc-user-operate/useroptservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	UserOptSvcRpcClient useroptservice.UserOptService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserOptSvcRpcClient: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserOptServiceConf)),
	}
}
