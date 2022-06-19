package logic

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-video-service/internal/svc"
	"douyin/service/rpc-video-service/videoSvcPb"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type FeedVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedVideosLogic {
	return &FeedVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedVideosLogic) FeedVideos(in *videoSvcPb.FeedVideosReq) (*videoSvcPb.FeedVideosResp, error) {

	// 转换为数据库里的时间格式
	reqTime := time.Unix(in.LastTime, 0)

	logx.Errorf("小于这个时间的视频reqTime: %v", reqTime)

	whereBuilder := l.svcCtx.VideoModel.RowBuilder().Where("create_time < ?", reqTime)
	list, err := l.svcCtx.VideoModel.FindPageListByPage(l.ctx, whereBuilder, 1, 10, "create_time DESC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get video list fail"), "get video list fail FindPageListByIdDESC  err : %v , lastTime:%d", err, in.LastTime)
	}
	/*
		如果想重构这里可以用泛型
	*/
	var videos []*videoSvcPb.Video
	videos = make([]*videoSvcPb.Video, 0, len(list))
	var NextTime int64
	if len(list) > 0 {
		for _, v := range list {
			var video videoSvcPb.Video
			_ = copier.Copy(&video, v)
			videos = append(videos, &video)
		}
		// 最早的视频的时间 时间戳最小的
		NextTime = list[len(list)-1].CreateTime.Unix()
	}

	logx.Errorf("下次请求的视频小于这个时间的视频: %v", list[len(list)-1].CreateTime)

	return &videoSvcPb.FeedVideosResp{
		VideoPubList: videos,
		NextTime:     NextTime,
	}, nil
}
