package logic

import (
	"context"
	"douyin/common/help/int64ToStr"
	"douyin/common/xerr"
	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
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

	whereBuilder := l.svcCtx.UserFollow.RowBuilder().Where("user_id = ? and In (?)", in.UserId, int64ToStr.Int64ToStr(in.AuthIds))
	res, err := l.svcCtx.UserFollow.FindAll(l.ctx, whereBuilder, "ASC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserFollow  err , id:%d , err:%v", in.UserId, err)
	}

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
