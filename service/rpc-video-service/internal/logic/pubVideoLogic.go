package logic

import (
	"context"
	"douyin/service/rpc-video-service/internal/svc"
	"douyin/service/rpc-video-service/model"
	"douyin/service/rpc-video-service/videoSvcPb"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PubVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPubVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubVideoLogic {
	return &PubVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------video-----------------------
func (l *PubVideoLogic) PubVideo(in *videoSvcPb.PubVideoReq) (*videoSvcPb.PubVideoResp, error) {

	_, err := l.svcCtx.VideoModel.Insert(l.ctx, nil, &model.Video{
		AuthId:   in.AuthId,
		Title:    in.Title,
		PlayURL:  in.PlayURL,
		CoverURL: in.CoverURL,
	})
	if err != nil {
		return nil, errors.Wrapf(err, " PubVideo insert video fail")
	}
	return &videoSvcPb.PubVideoResp{}, nil
}
