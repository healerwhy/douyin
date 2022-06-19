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

type GetUserFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavoriteLogic {
	return &GetUserFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFavorite 获取用户对视频的点赞
func (l *GetUserFavoriteLogic) GetUserFavorite(in *userOptPb.GetUserFavoriteReq) (*userOptPb.GetUserFavoriteResp, error) {

	if in.VideoIds != nil { // 给 feed 接口使用 查看用户对以下videoids的点赞状态
		whereBuilder := l.svcCtx.UserFavoriteModel.RowBuilder().Where(squirrel.Eq{"user_id": in.UserId, "video_id": in.VideoIds}).Where(squirrel.NotEq{"is_favorite": 0})

		res, err := l.svcCtx.UserFavoriteModel.FindAll(l.ctx, whereBuilder, "")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserFavorite  err , id:%d , err:%v", in.UserId, err)
		}

		// 该用户对视频的点赞映射
		var respUserFavoriteList map[int64]bool
		if len(res) > 0 {
			respUserFavoriteList = make(map[int64]bool, len(res))
			for _, v := range res {
				respUserFavoriteList[v.VideoId] = true
			}
		}

		return &userOptPb.GetUserFavoriteResp{
			UserFavoriteList: respUserFavoriteList,
		}, nil
	} else { // 给拉取点赞列表使用 查询的是用户对哪些视频进行了点赞
		whereBuilder := l.svcCtx.UserFavoriteModel.RowBuilder().Where("user_id = ? and is_favorite != 0 ", in.UserId)

		res, err := l.svcCtx.UserFavoriteModel.FindAll(l.ctx, whereBuilder, "user_id ASC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserFavorite  err , id:%d , err:%v", in.UserId, err)
		}

		/*
			这里用数组是因为拉取点赞列表的接口返回的是数组 可以直接返回 比较方便
		*/
		var respUserFavoriteArr []int64
		if len(res) > 0 {
			respUserFavoriteArr = make([]int64, len(res))
			for _, v := range res {
				respUserFavoriteArr = append(respUserFavoriteArr, v.VideoId)
			}
		}
		return &userOptPb.GetUserFavoriteResp{
			UserFavoriteArr: respUserFavoriteArr,
		}, nil
	}
}
