package feed

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/rpc-user-info/userInfoPb"
	"douyin/service/rpc-user-operate/userOptPb"
	"douyin/service/rpc-user-operate/useroptservice"
	"douyin/service/rpc-video-service/videoSvcPb"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"sort"
	"time"

	"douyin/service/api/internal/types"
)

type FeedVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedVideoListLogic {
	return &FeedVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func getAuthIdsAndVideoIds(VideoPubList []*videoSvcPb.Video) (authIds []int64, videoIds []int64) {
	// 把视频里的所用用户的信息找出来 并去重
	authIdsTemp := make(map[int64]interface{}, len(VideoPubList))
	authIds = make([]int64, 0, len(VideoPubList))
	// 构建videoIds  查询当前id对这些视频是否点赞
	videoIds = make([]int64, 0, len(VideoPubList))
	for _, v := range VideoPubList {
		if _, ok := authIdsTemp[v.AuthId]; !ok {
			authIdsTemp[v.AuthId] = nil
			authIds = append(authIds, v.AuthId)
		}
		videoIds = append(videoIds, v.Id)
	}
	sort.Slice(videoIds, func(i, j int) bool { return videoIds[i] < videoIds[j] }) // 升序排列
	sort.Slice(authIds, func(i, j int) bool { return authIds[i] < authIds[j] })    // 升序排列
	return authIds, videoIds
}

func (l *FeedVideoListLogic) FeedVideoList(req *types.FeedVideoListReq) (resp *types.FeedVideoListRes, err error) {
	// 查看LastTime是否存在 不存在就直接从video里的 拉取最多30条视频
	var LastTime int64
	if req.LastTime != 0 {
		LastTime = time.Unix(req.LastTime, 0).Unix()
	} else {
		LastTime = time.Now().Unix()
	}

	videos, err := l.svcCtx.VideoSvcRpcClient.FeedVideos(l.ctx, &videoSvcPb.FeedVideosReq{
		LastTime: LastTime,
	})
	if err != nil {
		logx.Errorf("get video list fail %s", err.Error())
		return &types.FeedVideoListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get video list fail",
			},
		}, nil
	}

	/*
		如果存在作者或者说有视频 才回去关注是否对视频点赞 对作者关注
	*/
	var videoList []*types.PubVideo // 最终返回的视频列表

	if len(videos.VideoPubList) > 0 {
		authIds, videoIds := getAuthIdsAndVideoIds(videos.VideoPubList)

		// 查询所有视频的的作者信息
		authsInfo, err := l.svcCtx.UserInfoRpcClient.AuthsInfo(l.ctx, &userInfoPb.AuthsInfoReq{ // 返回作者信息 按照作者id升序排列
			AuthIds: authIds,
		})
		if err != nil {
			return &types.FeedVideoListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "get authors info fail " + err.Error(),
				},
			}, nil
		}

		if req.Token != "" { // 因为带有token是登录状态所以有点赞状态以及关注状态
			userId := l.ctx.Value(myToken.CurrentUserId("LoginUserId")).(int64)

			var userFavoriteList *userOptPb.GetUserFavoriteResp
			var userFollowList *userOptPb.GetUserFollowResp
			var resp *types.FeedVideoListRes
			err := mr.Finish(func() error {
				userFavoriteList, err = l.svcCtx.UserOptSvcRpcClient.GetUserFavorite(l.ctx, &useroptservice.GetUserFavoriteReq{
					UserId:   userId,
					VideoIds: videoIds,
				})

				if err != nil {
					resp = &types.FeedVideoListRes{
						Status: types.Status{
							Code: xerr.ERR,
							Msg:  "get user favorite video relation fail " + err.Error(),
						},
					}
					return err
				}
				return nil

			}, func() error {
				// 构建following map 查询对这些作者是否关注
				userFollowList, err = l.svcCtx.UserOptSvcRpcClient.GetUserFollow(l.ctx, &useroptservice.GetUserFollowReq{
					UserId:  userId,
					AuthIds: authIds,
				})
				if err != nil {
					resp = &types.FeedVideoListRes{
						Status: types.Status{
							Code: xerr.ERR,
							Msg:  "get user follow author relation fail " + err.Error(),
						},
					}
					return err
				}
				return nil
			})
			if err != nil {
				return resp, nil
			}

			for _, v := range videos.VideoPubList {
				var video types.PubVideo
				_ = copier.Copy(&video, v)
				_ = copier.Copy(&video.Author, authsInfo.Auths[v.AuthId])

				// 用户对该视频是否点赞
				video.IsFavorite = userFavoriteList.UserFavoriteList[v.Id]
				// 用户对该视频的作者是否关注
				video.Author.IsFollow = userFollowList.UserFollowList[v.AuthId]

				videoList = append(videoList, &video)
			}
		} else { // 未登录直接把视频流返回
			for _, v := range videos.VideoPubList {
				var video types.PubVideo
				_ = copier.Copy(&video, v)
				_ = copier.Copy(&video.Author, authsInfo.Auths[v.AuthId])

				logx.Errorf("video %+v", video)
				video.IsFavorite = false
				video.Author.IsFollow = false

				videoList = append(videoList, &video)
			}
		}
	}

	return &types.FeedVideoListRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		VideoList: videoList,
		NextTime:  videos.NextTime,
	}, nil
}
