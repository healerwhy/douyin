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
type UserFavoriteOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFavoriteUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFavoriteOpt {
	return &UserFavoriteOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavoriteOpt) Consume(_, val string) error {
	var message types.UserFavoriteOptMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

// 处理逻辑
func (l *UserFavoriteOpt) execService(message types.UserFavoriteOptMessage) error {

	status := l.getUserOpt(message.OptStatus)
	if status != -99 {
		fmt.Printf("status: %d, %s \n", status, message.Opt)
	}

	return nil
}

//Get order status based on payment status.
func (l *UserFavoriteOpt) getUserOpt(Status int64) int64 {

	switch Status {
	case 1:
		return 1111
	case 2:
		return 2222
	default:
		return -99
	}

}
