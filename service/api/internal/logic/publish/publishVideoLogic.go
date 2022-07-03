package publish

import (
	"context"
	"douyin/common/help/cos"
	myToken "douyin/common/help/token"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"douyin/service/rpc-video-service/videoservice"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PubVideoReq) (resp *types.PubVideoRes, err error) {
	// 从前端获取视频
	file, _, err := l.r.FormFile("data")
	authId := l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)

	upLoader := cos.UploaderVideo{
		UserId:      authId,
		MachineId:   l.svcCtx.Config.COSConf.MachineId,
		VideoBucket: l.svcCtx.Config.COSConf.VideoBucket,
		SecretID:    l.svcCtx.Config.COSConf.SecretId,
		SecretKey:   l.svcCtx.Config.COSConf.SecretKey,
	}
	key, err := upLoader.UploadVideo(l.ctx, file)
	if err != nil {
		logx.Errorf("upload video failed: %s", err.Error())
		return &types.PubVideoRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "upload video failed",
			},
		}, nil
	}

	// 将视频信息存到数据库
	_, err = l.svcCtx.VideoSvcRpcClient.PubVideo(l.ctx, &videoservice.PubVideoReq{
		AuthId:   authId,
		Title:    req.Title,
		PlayURL:  l.svcCtx.Config.COSConf.VideoBucket + "/" + key + ".mp4",
		CoverURL: l.svcCtx.Config.COSConf.CoverBucket + "/" + key + "_0.jpg",
	})
	if err != nil {
		logx.Errorf("publish video failed: %s", err.Error())
		return &types.PubVideoRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "server database err",
			},
		}, nil
	}

	return &types.PubVideoRes{
		Status: types.Status{
			Code: xerr.OK,
		},
	}, nil
}
