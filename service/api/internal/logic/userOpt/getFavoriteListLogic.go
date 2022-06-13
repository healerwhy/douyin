package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-user-info/userInfoPb"
	"douyin/service/rpc-user-operate/useroptservice"
	"douyin/service/rpc-video-service/videoSvcPb"
	"github.com/jinzhu/copier"

	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteListLogic {
	return &GetFavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFavoriteListLogic) GetFavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListRes, err error) {
	videosId, err := l.svcCtx.UserOptSvcRpcClient.GetUserFavorite(l.ctx, &useroptservice.GetUserFavoriteReq{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.FavoriteListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get user favorite list fail " + err.Error(),
			},
		}, nil
	}

	var videoList []*types.PubVideo // 最终返回的视频列表

	// 存在点赞的视频
	if videosId != nil {
		// 拿着这些视频id去查询视频信息
		videoArr, err := l.svcCtx.VideoSvcRpcClient.GetMyFavoriteVideos(l.ctx, &videoSvcPb.MyFavoriteVideosReq{
			VideoIdArr: videosId.UserFavoriteArr,
		})

		if err != nil {
			return &types.FavoriteListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "get video list fail " + err.Error(),
				},
			}, nil
		}

		// 把视频里的作者的信息找出来 并去重
		var authIdsTemp map[int64]interface{}
		var authIds []int64
		authIdsTemp = make(map[int64]interface{}, len(videoArr.VideoPubList))
		for _, v := range videoArr.VideoPubList {
			authIdsTemp[v.AuthId] = nil
		}
		authIds = make([]int64, 0, len(authIdsTemp))
		for k := range authIdsTemp {
			authIds = append(authIds, k)
		}

		// 查询所有视频的的作者信息
		authsInfo, err := l.svcCtx.UserInfoRpcClient.AuthsInfo(l.ctx, &userInfoPb.AuthsInfoReq{ // 返回作者信息 按照作者id升序排列
			AuthIds: authIds,
		})

		if err != nil {
			return &types.FavoriteListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "get author list fail " + err.Error(),
				},
			}, nil
		}

		// 查找当前用户对作者的关注状态
		userFollowList, err := l.svcCtx.UserOptSvcRpcClient.GetUserFollow(l.ctx, &useroptservice.GetUserFollowReq{
			UserId:  req.UserId,
			AuthIds: authIds,
		})

		for _, v := range videoArr.VideoPubList {
			var video types.PubVideo
			_ = copier.Copy(&video, v)
			_ = copier.Copy(&video.Author, authsInfo.Auths[v.AuthId])
			// 点赞列表里肯定点赞了
			video.IsFavorite = true
			// 用户对该视频的作者是否关注
			video.Author.IsFollow = userFollowList.UserFollowList[v.AuthId]

			videoList = append(videoList, &video)
		}

		return &types.FavoriteListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			FavoriteList: videoList,
		}, nil
	} else { // 没有点赞的视频

		return &types.FavoriteListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			FavoriteList: nil,
		}, nil
	}

}
