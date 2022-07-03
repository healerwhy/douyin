package svc

import (
	"douyin/service/asynqJob/server/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Infof("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %s", err.Error())
				return true
			},
			Concurrency: 10, //max concurrent process server task num
		},
	)
}
