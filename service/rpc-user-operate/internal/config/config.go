package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	CacheConf cache.CacheConf

	COSConf struct {
		SecretId      string
		SecretKey     string
		CommentBucket string
	}
}
