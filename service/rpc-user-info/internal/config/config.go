package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisCacheConf redis.RedisConf
	DB             struct {
		DataSource string
	}
	CacheConf cache.CacheConf
}
