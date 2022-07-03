package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// KqConfig 用于发现kafka的Topic
type KqConfig struct {
	Brokers []string
	Topic   string
}

type Config struct {
	rest.RestConf

	Cache cache.CacheConf
	// rpc
	UserInfoService zrpc.RpcClientConf
	VideoService    zrpc.RpcClientConf
	UserOptService  zrpc.RpcClientConf
	// 腾讯云
	COSConf struct {
		SecretId    string
		SecretKey   string
		MachineId   uint16
		VideoBucket string
		CoverBucket string
	}
	// kafka
	UserFavoriteOptServiceConf KqConfig
	UserCommentOptServiceConf  KqConfig
	UserFollowOptServiceConf   KqConfig
}
