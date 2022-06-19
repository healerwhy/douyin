package logic

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-video-service/model"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
	whereBuilder := l.svcCtx.VideoModel.RowBuilder().Where(squirrel.Eq{"id": in.VideoIdArr})
	list, err := l.svcCtx.VideoModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get video list err : %v , in :%+v", err, in)
	}

	/*
		如果想重构这里可以用泛型
	*/
	var resp []*videoSvcPb.Video
	if len(list) > 0 {
		for _, video := range list {
			var pbVideo videoSvcPb.Video
			_ = copier.Copy(&pbVideo, video)

			resp = append(resp, &pbVideo)
		}
	}

	return &videoSvcPb.MyFavoriteVideosResp{
		VideoPubList: resp,
	}, nil
}
