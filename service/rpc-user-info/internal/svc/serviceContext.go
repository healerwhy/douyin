package svc

import (
	"douyin/service/rpc-user-info/internal/config"
	"douyin/service/rpc-user-info/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	RedisCache *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		RedisCache: c.RedisCacheConf.NewRedis(),
		UserModel:  model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
