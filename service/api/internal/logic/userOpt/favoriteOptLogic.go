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

	// 前端传入的是1，2表示点赞与取消点赞，入口这里就将它转换成1，0表示点赞与取消点赞
	msgTemp.ActionType = l.getActionType(req.ActionType)
	if msgTemp.ActionType == -99 {
		logx.Errorf("error actionType : %d", req.ActionType)
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "operate error",
			},
		}, nil
	}

	msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)

	// 序列化
	msg, err := json.Marshal(msgTemp)
	if err != nil {
		logx.Errorf("FavoriteOpt json.Marshal err : %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), " json.Marshal err")
	}

	// 向消息队列发送消息
	err = l.svcCtx.FavoriteOptMsgProducer.Push(string(msg))
	if err != nil {
		logx.Errorf("FavoriteOpt msgProducer.Push err : %s", err.Error())
		return &types.FavoriteOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to FavoriteOptMsgConsumer err",
			},
		}, nil
	}

	return &types.FavoriteOptRes{
		Status: types.Status{
			Code: xerr.OK,
		},
	}, nil
}

func (l *FavoriteOptLogic) getActionType(actionType int64) int64 {

	switch actionType { // 方便扩展
	case messageTypes.ActionADD:
		return 1
	case messageTypes.ActionCancel:
		return 0
	default:
		return -99
	}
}
