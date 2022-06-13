package logic

import (
	"context"

	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFavoriteLogic {
	return &AddFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userFavoriteList-----------------------
func (l *AddFavoriteLogic) AddFavorite(in *userOptPb.AddFavoriteReq) (*userOptPb.AddFavoriteResp, error) {
	// todo: add your logic here and delete this line

	return &userOptPb.AddFavoriteResp{}, nil
}
