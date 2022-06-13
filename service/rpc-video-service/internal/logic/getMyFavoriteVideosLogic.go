package logic

import (
	"context"

	"douyin/service/rpc-video-service/internal/svc"
	"douyin/service/rpc-video-service/videoSvcPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyFavoriteVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyFavoriteVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFavoriteVideosLogic {
	return &GetMyFavoriteVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyFavoriteVideosLogic) GetMyFavoriteVideos(in *videoSvcPb.MyFavoriteVideosReq) (*videoSvcPb.MyFavoriteVideosResp, error) {
	// todo: add your logic here and delete this line

	return &videoSvcPb.MyFavoriteVideosResp{}, nil
}
