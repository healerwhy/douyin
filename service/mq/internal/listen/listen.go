package listen

import (
	"context"
	"douyin/service/mq/internal/config"
	"douyin/service/mq/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/service"
)

const scriptLoadSADD = "redis.call('SADD', KEYS[1], ARGV[1]) redis.call('SADD', KEYS[2], ARGV[2]) redis.call('EXPIRE', KEYS[1], 60) redis.call('EXPIRE', KEYS[2], 60)"

// Mqs back to all consumers
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	// 加载脚本
	tmp, err := svcContext.RedisCache.ScriptLoadCtx(ctx, scriptLoadSADD)
	if err != nil {
		logx.Errorf("load script err:%+v", err)
		return nil
	}
	svcContext.ScriptADD = tmp

	var services []service.Service

	//kq ：pub sub
	services = append(services, KqMqs(c, ctx, svcContext)...)

	return services
}
