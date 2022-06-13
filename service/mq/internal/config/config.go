package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	// kq : pub sub
	// 点赞
	UserFavoriteOptServiceConf kq.KqConf
	// 评论
	UserCommentOptServiceConf kq.KqConf
	// 关注
	UserFollowOptServiceConf kq.KqConf

	// rpc
	UserOptService zrpc.RpcClientConf
}
