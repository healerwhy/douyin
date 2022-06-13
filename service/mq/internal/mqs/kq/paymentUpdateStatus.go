package kq

import (
	"context"
	"douyin/service/mq/internal/kqueue"
	"douyin/service/mq/internal/svc"
	"douyin/service/rpc-user-operate/useroptservice"
	"encoding/json"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
	Listening to the payment flow status change notification message queue
*/
type UserFavoriteOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentUpdateStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFavoriteOpt {
	return &UserFavoriteOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavoriteOpt) Consume(_, val string) error {

	var message kqueue.UserFavoriteOptMessage
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
func (l *UserFavoriteOpt) execService(message kqueue.UserFavoriteOptMessage) error {

	orderTradeState := l.getUserOpt(message.OptStatus)
	if orderTradeState != -99 {
		//update homestay order state
		_, err := l.svcCtx.UserOptSvcRpcClient.AddFavorite(l.ctx, &useroptservice.AddFavoriteReq{})
		if err != nil {
			return errors.Wrap(err, " add favorite fail")
		}
	}

	return nil
}

//Get order status based on payment status.
func (l *UserFavoriteOpt) getUserOpt(Status int64) int64 {

	switch 1 {
	case 2:
		return 2
	case 3:
		return 1
	default:
		return -99
	}

}
