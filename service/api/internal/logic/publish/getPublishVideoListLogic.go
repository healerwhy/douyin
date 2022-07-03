package publish

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"douyin/service/rpc-user-info/userInfoPb"
	"douyin/service/rpc-video-service/videoservice"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPublishVideoListLogic) GetPublishVideoList(req *types.GetPubVideoListReq) (resp *types.GetPubVideoListRes, err error) {
	// 获得本人发布列表
	res, err := l.svcCtx.VideoSvcRpcClient.GetVideoList(l.ctx, &videoservice.GetVideoListReq{
		AuthId: req.UserId,
	})
	if err != nil {
		logx.Errorf("get publish video list fail %s", err.Error())
		return &types.GetPubVideoListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get publish video list failed",
			},
		}, nil
	}

	// 如果没有视频直接返回
	if len(res.VideoPubList) == 0 {
		return &types.GetPubVideoListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
		}, nil
	}

	// 获得本人的信息
	authInfo, err := l.svcCtx.UserInfoRpcClient.Info(l.ctx, &userInfoPb.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.GetPubVideoListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get user info failed",
			},
		}, nil
	}

	// 整合数据
	// 将rpc返回的videoPubList转换为api返回的videoPubList
	var videos []*types.PubVideo
	for _, v := range res.VideoPubList {
		var video types.PubVideo
		_ = copier.Copy(&video, v)
		_ = copier.Copy(&video.Author, authInfo.User)

		videos = append(videos, &video)
	}

	return &types.GetPubVideoListRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		VideoPubList: videos,
	}, nil
}
