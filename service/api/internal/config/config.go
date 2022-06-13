package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//JWTAuth struct {
	//	AccessSecret string
	//	AccessExpire int64
	//}
	Cache           cache.CacheConf
	UserInfoService zrpc.RpcClientConf
	VideoService    zrpc.RpcClientConf
	UserOptService  zrpc.RpcClientConf

	COSConf struct {
		SecretId    string
		SecretKey   string
		MachineId   uint16
		VideoBucket string
		CoverBucket string
	}
}
