package logic

import (
	"context"
	"douyin/service/asynqJob/server/internal/logic/jobs"
	"douyin/service/asynqJob/server/internal/svc"
	"douyin/service/asynqJob/server/jobtype"
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

	// handle
	mux.Handle(jobtype.ScheduleGetUserFavoriteStatus, jobs.NewGetUserFavoriteStatusHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleGetUserFollowStatus, jobs.NewGetUserFollowStatusHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleGetUserComment, jobs.NewGetUserCommentHandler(l.svcCtx))

	return mux
}
