package userOpt

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/messageTypes"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteOptLogic {
	return &FavoriteOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteOptLogic) FavoriteOpt(req *types.FavoriteOptReq) (resp *types.FavoriteOptRes, err error) {
	var msgTemp messageTypes.UserFavoriteOptMessage
	_ = copier.Copy(&msgTemp, req)
	msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)

	logx.Errorf("FavoriteOpt msgTemp : %+v", msgTemp)

	// 序列化
	msg, err := json.Marshal(msgTemp)
	if err != nil {
		return nil, errors.Wrapf(err, " json.Marshal err")
	}
	// 向消息队列发送消息
	err = l.svcCtx.FavoriteOptMsgProducer.Push(string(msg))
	if err != nil {
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to favoriteOptMsgConsumer err",
			},
		}, nil
	}
	return &types.FavoriteOptRes{
		Status: types.Status{
			Code: xerr.OK,
			Msg:  "operate success",
		},
	}, nil
}
