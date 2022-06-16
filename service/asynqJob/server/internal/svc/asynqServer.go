package svc

import (
	"douyin/service/asynqTask/server/internal/config"
	"fmt"
	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 10, //max concurrent process server task num
		},
	)
}
