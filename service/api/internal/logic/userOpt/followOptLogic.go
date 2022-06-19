package userOpt

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/messageTypes"
	"douyin/common/xerr"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowOptLogic {
	return &FollowOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowOptLogic) FollowOpt(req *types.FollowOptReq) (resp *types.FollowOptRes, err error) {
	var msgTemp messageTypes.UserFollowOptMessage
	_ = copier.Copy(&msgTemp, req)

	// 前端传入的是1，2表示关注与取消关注，入口这里就将它转换成1，0表示点赞与取消关注
	msgTemp.ActionType = l.getActionType(req.ActionType)

	if msgTemp.ActionType == -99 {
		return &types.FollowOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "operate error",
			},
		}, nil
	}
	msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
	logx.Infof("FollowOpt msgTemp : %+v", msgTemp)

	// 序列化
	msg, err := json.Marshal(msgTemp)
	if err != nil {
		return nil, errors.Wrapf(err, " json.Marshal err")
	}

	// 向消息队列发送消息
	err = l.svcCtx.FollowOptMsgProducer.Push(string(msg))
	if err != nil {
		return &types.FollowOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to FollowOptMsgProducer err",
			},
		}, nil
	}

	return &types.FollowOptRes{
		Status: types.Status{
			Code: xerr.OK,
			Msg:  "operate success",
		},
	}, nil
}

func (l *FollowOptLogic) getActionType(actionType int64) int64 {

	switch actionType { // 方便扩展
	case messageTypes.ActionADD:
		return 1
	case messageTypes.ActionCancel:
		return 0
	default:
		return -99
	}
}
