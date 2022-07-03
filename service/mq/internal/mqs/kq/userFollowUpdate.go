package kq

import (
	"context"
	"douyin/common/globalkey"
	"douyin/common/messageTypes"
	"douyin/service/mq/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

	logx.Infof("UserFollowOptMessage message : %+v\n", message)

	var req userOptPb.UpdateFollowStatusReq
	_ = copier.Copy(&req, &message)

	// 构造redis的数据
	dataKey := fmt.Sprintf(globalkey.FollowSetValTpl, message.FollowId)
	followSetVal := fmt.Sprintf(globalkey.FollowSetValTpl, message.FollowId)
	dataVal := fmt.Sprintf(globalkey.ExistDataValTpl, message.UserId, message.ActionType)

	// 消息取出来之后无非是点赞或者取消点赞 0，1，那么打到redis也是0，1
	_, err := l.svcCtx.RedisCache.EvalShaCtx(l.ctx, l.svcCtx.ScriptADD, []string{globalkey.FollowSetKey, dataKey}, []string{followSetVal, dataVal})
	if err != redis.Nil {
		logx.Errorf("script exec err : %v", err)
		return err
	}

	return nil
}
