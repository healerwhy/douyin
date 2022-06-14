package kq

import (
	"context"
	"douyin/service/mq/internal/svc"
	"douyin/service/mq/internal/types"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
	Listening to the payment flow status change notification message queue
*/
type UserCommentOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCommentUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserCommentOpt {
	return &UserCommentOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCommentOpt) Consume(_, val string) error {
	var message types.UserCommentOptMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("NewUserCommentUpdateMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("NewUserCommentUpdateMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

// 处理逻辑
func (l *UserCommentOpt) execService(message types.UserCommentOptMessage) error {

	status := l.getUserOpt(message.OptStatus)
	if status != -99 { //update mysql judging by status
		fmt.Printf("status: %d, %s \n", status, message.Opt)
	}

	return nil
}

//Get order status based on payment status.
func (l *UserCommentOpt) getUserOpt(Status int64) int64 {

	switch Status { //
	case 1:
		return 1111
	case 2:
		return 2222
	default:
		return -99
	}

}
