package logic

import (
	"context"
	"douyin/common/xerr"
	"github.com/pkg/errors"

	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowerLogic {
	return &GetUserFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowerLogic) GetUserFollower(in *userOptPb.GetUserFollowerReq) (*userOptPb.GetUserFollowerResp, error) {
	whereBuilder := l.svcCtx.UserFollowModel.RowBuilder().Where("follow_id = ? and is_follow != 0 ", in.UserId)

	res, err := l.svcCtx.UserFollowModel.FindAll(l.ctx, whereBuilder, "user_id ASC")
	if err != nil {
		logx.Errorf("GetUserFollowerLogic GetUserFollower err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserFollower  err , id:%d , err:%v", in.UserId, err)
	}

	/*
		这里用数组是因为拉取点赞列表的接口返回的是数组 可以直接返回 比较方便
	*/
	var userFollowerList map[int64]bool
	if len(res) > 0 {
		userFollowerList = make(map[int64]bool, len(res))
		for _, v := range res {
			userFollowerList[v.UserId] = true
		}
	}

	return &userOptPb.GetUserFollowerResp{
		UserFollowerList: userFollowerList,
	}, nil
}
