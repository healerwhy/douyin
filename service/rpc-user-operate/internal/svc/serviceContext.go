package svc

import (
	"douyin/service/rpc-user-operate/internal/config"
	"douyin/service/rpc-user-operate/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	UserFavorite model.UserFavoriteListModel
	UserFollow   model.UserFollowListModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		UserFavorite: model.NewUserFavoriteListModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		UserFollow:   model.NewUserFollowListModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
