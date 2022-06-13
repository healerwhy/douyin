package listen

import (
	"context"
	"douyin/service/mq/internal/config"
	kqMq "douyin/service/mq/internal/mqs/kq"
	"douyin/service/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

//pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.UserFavoriteOptServiceConf, kqMq.NewPaymentUpdateStatusMq(ctx, svcContext)),
		//.....
	}

}
