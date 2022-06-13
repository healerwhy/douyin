package logic

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-video-service/internal/svc"
	"douyin/service/rpc-video-service/model"
	"douyin/service/rpc-video-service/videoSvcPb"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListLogic) GetVideoList(in *videoSvcPb.GetVideoListReq) (*videoSvcPb.GetVideoListResp, error) {
	whereBuilder := l.svcCtx.VideoModel.RowBuilder().Where(squirrel.Eq{"auth_id": in.AuthId})
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

	return &videoSvcPb.GetVideoListResp{
		VideoPubList: resp,
	}, nil
}
