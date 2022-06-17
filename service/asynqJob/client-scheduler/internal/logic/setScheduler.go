package logic

import (
	"douyin/service/asynqJob/server/jobtype"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// GetUserFavoriteStatusScheduler 向Redis发送定时消息调用worker进行工作
func (l *MqueueScheduler) GetUserFavoriteStatusScheduler() {

	task := asynq.NewTask(jobtype.ScheduleGetUserFavoriteStatus, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("@every 20s", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【ScheduleGetUserFavoriteStatus】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【ScheduleGetUserFavoriteStatus】 registered an  entry: %q \n", entryID)
}

func (l *MqueueScheduler) GetUserFollowStatusScheduler() {

	task := asynq.NewTask(jobtype.ScheduleGetUserFollowStatus, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("@every 8m", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【ScheduleGetUserFollowStatus】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【ScheduleGetUserFollowStatus】 registered an  entry: %q \n", entryID)
}

func (l *MqueueScheduler) GetUserCommentScheduler() {

	task := asynq.NewTask(jobtype.ScheduleGetUserComment, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("@every 8m", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【ScheduleGetUserComment】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【ScheduleGetUserComment】 registered an  entry: %q \n", entryID)
}
