package svc

import (
	"douyin/service/rpc-video-service/internal/config"
	videoModel "douyin/service/rpc-video-service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel videoModel.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		VideoModel: videoModel.NewVideoModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
