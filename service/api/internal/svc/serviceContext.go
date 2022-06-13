package svc

import (
	"douyin/service/api/internal/config"
	"douyin/service/api/internal/middleware"
	"douyin/service/rpc-user-info/userinfoservice"
	"douyin/service/rpc-user-operate/useroptservice"
	"douyin/service/rpc-video-service/videoservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	UserInfoRpcClient   userinfoservice.UserInfoService
	VideoSvcRpcClient   videoservice.VideoService
	UserOptSvcRpcClient useroptservice.UserOptService
	AuthJWT             rest.Middleware
	IsLogin             rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserInfoRpcClient:   userinfoservice.NewUserInfoService(zrpc.MustNewClient(c.UserInfoService)),
		VideoSvcRpcClient:   videoservice.NewVideoService(zrpc.MustNewClient(c.VideoService)),
		UserOptSvcRpcClient: useroptservice.NewUserOptService(zrpc.MustNewClient(c.UserOptService)),
		AuthJWT:             middleware.NewAuthJWTMiddleware().Handle,
		IsLogin:             middleware.NewIsLoginMiddleware().Handle,
	}
}
