package userOpt

import (
	"context"
	"douyin/common/help/cos"
	myToken "douyin/common/help/token"
	"douyin/common/messageTypes"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"douyin/service/rpc-user-info/userInfoPb"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentOptLogic {
	return &CommentOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentOptLogic) CommentOpt(req *types.CommentOptReq) (resp *types.CommentOptRes, err error) {

	// 前端传入的是1，2表示评论取消评论，入口这里就将它转换成1，0表示评论取消评论
	msgTemp, status, err := l.getActionType(req)

	if msgTemp.ActionType == -99 || err != nil {
		return status, nil
	}
	if msgTemp.ActionType == 1 {
		// 拉取发布消息的用户信息
		userInfo, err := l.svcCtx.UserInfoRpcClient.Info(l.ctx, &userInfoPb.UserInfoReq{
			UserId: msgTemp.UserId,
		})
		if err != nil {
			return &types.CommentOptRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "get user info err",
				},
			}, nil
		}
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			Comment: &types.Comment{
				CommentId: msgTemp.CommentId,
				User: types.Author{
					UserId:        userInfo.User.UserId,
					UserName:      userInfo.User.UserName,
					FollowCount:   userInfo.User.FollowCount,
					FollowerCount: userInfo.User.FollowerCount,
					IsFollow:      false,
				},
				Content:    msgTemp.CommentText,
				CreateTime: msgTemp.CreateDate,
			},
		}, nil
	}

	if msgTemp.ActionType == 0 {
		return &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.OK,
			},
		}, nil
	}

	return nil, nil
}

func (l *CommentOptLogic) getActionType(req *types.CommentOptReq) (*messageTypes.UserCommentOptMessage, *types.CommentOptRes, error) {
	var msgTemp messageTypes.UserCommentOptMessage
	_ = copier.Copy(&msgTemp, req)

	switch req.ActionType { // 方便扩展
	case messageTypes.ActionADD:
		msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		msgTemp.CreateDate = time.Now().Format("01-02")
		msgTemp.ActionType = 1
		var genId cos.GenSnowFlake
		id, _ := genId.GenSnowFlake(1)
		msgTemp.CommentId = int64(id)

	case messageTypes.ActionCancel:
		msgTemp.UserId = l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
		msgTemp.ActionType = 0
	default:
		msgTemp.ActionType = -99
		return nil, &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer ActionType err",
			},
		}, errors.New("operate error")
	}

	logx.Infof("CommentOpt msgTemp : %+v", msgTemp)
	// 序列化
	msg, err := json.Marshal(msgTemp)
	if err != nil {
		return nil, &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer json.Marshal err",
			},
		}, errors.Wrapf(err, " json.Marshal err")
	}
	// 向消息队列发送消息
	err = l.svcCtx.CommentOptMsgProducer.Push(string(msg))
	if err != nil {
		return nil, &types.CommentOptRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "send message to CommentOptMsgConsumer err",
			},
		}, errors.Wrapf(err, " json.Marshal err")
	}
	return &msgTemp, nil, nil
}
