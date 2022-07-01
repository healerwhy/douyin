package listen

import (
	"context"
	"douyin/service/mq/internal/config"
	"douyin/service/mq/internal/svc"

	kqMq "douyin/service/mq/internal/mqs/kq"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// KqMqs pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.UserFavoriteOptServiceConf, kqMq.NewUserFavoriteUpdateMq(ctx, svcContext)),
		kq.MustNewQueue(c.UserFollowOptServiceConf, kqMq.NewUserFollowUpdateMq(ctx, svcContext)),
		kq.MustNewQueue(c.UserCommentOptServiceConf, kqMq.NewUserCommentUpdateMq(ctx, svcContext)),
		// 配置
		//.....
	}

}
