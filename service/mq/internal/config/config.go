package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	// kq : pub sub
	// 点赞 评论 关注
	UserFavoriteOptServiceConf kq.KqConf
	UserCommentOptServiceConf  kq.KqConf
	UserFollowOptServiceConf   kq.KqConf
}
