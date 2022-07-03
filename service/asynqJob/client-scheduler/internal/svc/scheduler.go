package svc

import (
	"douyin/service/asynqJob/client-scheduler/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// create client-scheduler
func newScheduler(c config.Config) *asynq.Scheduler {

	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				logx.Infof("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
			},
		})
}
