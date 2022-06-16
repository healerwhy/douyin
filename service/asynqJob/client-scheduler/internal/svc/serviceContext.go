package svc

import (
	"douyin/service/asynqJob/client-scheduler/internal/config"
	"github.com/hibiken/asynq"
)

type ServiceContext struct {
	Config config.Config

	Scheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Scheduler: newScheduler(c),
	}
}
