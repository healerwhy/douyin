package kq

import (
	"context"
	"douyin/common/messageTypes"
	"douyin/service/mq/internal/svc"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
	Listening to the payment flow status change notification message queue
*/
type UserFollowOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFollowUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowOpt {
	return &UserFollowOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFollowOpt) Consume(_, val string) error {
	var message messageTypes.UserFollowOptMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserFollowOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserFollowOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

// 处理逻辑
func (l *UserFollowOpt) execService(message messageTypes.UserFollowOptMessage) error {

	//status := l.getUserOpt(message.ActionType)
	//if status != -99 {
	//	fmt.Printf("status: %d, %s \n", status, message.VideoId)
	//}

	return nil
}
