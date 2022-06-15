package kq

import (
	"context"
	"douyin/common/messageTypes"
	"douyin/service/mq/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
	"encoding/json"
	"github.com/jinzhu/copier"
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

func NewUserFavoriteUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFavoriteOpt {
	return &UserFavoriteOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavoriteOpt) Consume(_, val string) error {
	var message messageTypes.UserFavoriteOptMessage

	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserFavoriteOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserFavoriteOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}
	return nil

}

// 处理逻辑
func (l *UserFavoriteOpt) execService(message messageTypes.UserFavoriteOptMessage) error {
	logx.Infof("UserFavoriteOptMessage message : %+v\n", message)

	status := l.getUserOpt(message.ActionType)
	if status == messageTypes.ActionErr {
		return errors.New("unknown action type")
	}
	var req userOptPb.UpdateFavoriteStatusReq
	_ = copier.Copy(&req, &message)
	req.ActionType = status
	_, err := l.svcCtx.UserOptSvcRpcClient.UpdateFavoriteStatus(l.ctx, &req)
	if err != nil {
		return err
	}
	return nil

}

//Get order status based on payment status.
func (l *UserFavoriteOpt) getUserOpt(Status int64) int64 {

	switch Status { // 1:add, 2:cancel 这样就可以扩展状态 而不用if 吧啦吧啦 not to be shit
	case messageTypes.ActionADD: // 那么让is_favorite = 0
		return 1
	case messageTypes.ActionCancel: //
		return 0
	default:
		return messageTypes.ActionErr
	}

}
