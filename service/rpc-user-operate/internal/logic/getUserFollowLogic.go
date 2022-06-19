package logic

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowLogic {
	return &GetUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowLogic) GetUserFollow(in *userOptPb.GetUserFollowReq) (*userOptPb.GetUserFollowResp, error) {

	if in.AuthIds != nil {
		whereBuilder := l.svcCtx.UserFollowModel.RowBuilder().Where(squirrel.Eq{"user_id": in.UserId, "follow_id": in.AuthIds}).Where(squirrel.NotEq{"is_follow": 0})
		res, err := l.svcCtx.UserFollowModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserFollow  err , id:%d , err:%v", in.UserId, err)
		}

		//logx.Errorf("GetUserFollow  %#v", res)

		var userFollowList map[int64]bool
		if len(res) > 0 {
			userFollowList = make(map[int64]bool, len(res))
			for _, v := range res {
				userFollowList[v.UserId] = true
			}
		}

		return &userOptPb.GetUserFollowResp{
			UserFollowList: userFollowList,
		}, nil

	} else {
		whereBuilder := l.svcCtx.UserFollowModel.RowBuilder().Where("user_id = ? and is_follow != 0 ", in.UserId)

		res, err := l.svcCtx.UserFollowModel.FindAll(l.ctx, whereBuilder, "user_id ASC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserFollow  err , id:%d , err:%v", in.UserId, err)
		}

		/*
			这里用数组是因为拉取点赞列表的接口返回的是数组 可以直接返回 比较方便
		*/
		var userFollowList map[int64]bool
		if len(res) > 0 {
			userFollowList = make(map[int64]bool, len(res))
			for _, v := range res {
				userFollowList[v.UserId] = true
			}
		}

		return &userOptPb.GetUserFollowResp{
			UserFollowList: userFollowList,
		}, nil

	}

}
