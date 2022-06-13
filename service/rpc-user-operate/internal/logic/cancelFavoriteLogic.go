package logic

import (
	"context"

	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelFavoriteLogic {
	return &CancelFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelFavoriteLogic) CancelFavorite(in *userOptPb.CancelFavoriteResp) (*userOptPb.CancelFavoriteResp, error) {
	// todo: add your logic here and delete this line

	return &userOptPb.CancelFavoriteResp{}, nil
}
