package svc

import (
	usermodel "douyin/service/rpc-user-info/model"
	"douyin/service/rpc-user-operate/internal/config"
	"douyin/service/rpc-user-operate/model"
	videomodel "douyin/service/rpc-video-service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	UserFavoriteModel model.UserFavoriteListModel
	UserFollowModel   model.UserFollowListModel
	UserCommentModel  model.UserCommentListModel
	UserModel         usermodel.UserModel
	VideoModel        videomodel.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserFavoriteModel: model.NewUserFavoriteListModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		UserFollowModel:   model.NewUserFollowListModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		UserModel:         usermodel.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		VideoModel:        videomodel.NewVideoModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
		UserCommentModel:  model.NewUserCommentListModel(sqlx.NewMysql(c.DB.DataSource), c.CacheConf),
	}
}
