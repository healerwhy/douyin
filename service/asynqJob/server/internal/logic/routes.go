package logic

import (
	"context"
	"douyin/service/asynqTask/server/internal/logic/tasks"
	"douyin/service/asynqTask/server/internal/svc"
	"douyin/service/asynqTask/server/jobtype"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register server
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//client-scheduler server
	mux.Handle(jobtype.ScheduleGetUserFavoriteStatus, tasks.NewGetUserFavoriteStatusHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleGetUserFollowStatus, tasks.NewGetUserFollowStatusHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleGetUserComment, tasks.NewGetUserCommentHandler(l.svcCtx))

	return mux
}
