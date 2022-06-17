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
		logx.Errorf("UserFavoriteOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}
	return nil
}

// 处理逻辑
func (l *UserFavoriteOpt) execService(message messageTypes.UserFavoriteOptMessage) error {

	logx.Infof("UserFavoriteOptMessage message : %+v\n", message)

	var req userOptPb.UpdateFavoriteStatusReq
	_ = copier.Copy(&req, &message)

	// 构造redis的数据
	dataKey := fmt.Sprintf(globalkey.FavoriteTpl, message.VideoId)
	favoriteSetVal := fmt.Sprintf(globalkey.FavoriteTpl, message.VideoId)
	dataVal := fmt.Sprintf(globalkey.DataValTpl, message.UserId, message.ActionType)

	_, err := l.svcCtx.RedisCache.EvalShaCtx(l.ctx, l.svcCtx.ScriptTag, []string{globalkey.FavoriteSetKey, dataKey}, []string{favoriteSetVal, dataVal})
	if err != redis.Nil {
		logx.Errorf("script exec err : %v", err)
		return err
	}

	return nil
}
