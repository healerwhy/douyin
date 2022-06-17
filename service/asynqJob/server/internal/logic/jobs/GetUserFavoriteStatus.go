package jobs

import (
	"context"
	"douyin/common/globalkey"
	"douyin/service/asynqJob/server/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"strconv"
	"strings"
)

// GetUserFavoriteStatusHandler   shcedule billing to home business
type GetUserFavoriteStatusHandler struct {
	svcCtx *svc.ServiceContext
}

func NewGetUserFavoriteStatusHandler(svcCtx *svc.ServiceContext) *GetUserFavoriteStatusHandler {
	return &GetUserFavoriteStatusHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask every one minute exec : if return err != nil , asynq will retry
func (l *GetUserFavoriteStatusHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	logx.Infof("NewGetUserFavoriteStatusHandler server -----> every 5 minutes exec ")
	vals, err := l.svcCtx.RedisCache.SmembersCtx(ctx, globalkey.FavoriteSetKey)
	if err != nil {
		logx.Errorf("RedisCache.SmembersCtx error -----> %v", err)
		return err
	}
	if len(vals) == 0 {
		logx.Infof("RedisCache.SmembersCtx no data")
		return nil
	}
	mr.ForEach(func(source chan<- interface{}) {
		for _, videoIdKey := range vals {
			source <- videoIdKey
		}
	}, func(item interface{}) {
		videoIdKey := item.(string)
		users, err := l.svcCtx.RedisCache.SmembersCtx(ctx, videoIdKey)
		if err != nil { // 获取赞了这个视频的所有的用户Id
			logx.Errorf("RedisCache.SmembersCtx error -----> %v", err)
			return
		}
		// 切分出视频的Id
		_, videoIdStr, _ := strings.Cut(videoIdKey, ":")
		videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

		mr.ForEach(func(source chan<- interface{}) {
			for _, userId := range users {
				source <- userId
			}
		}, func(item interface{}) {
			// 切分出用户的Id 和 点赞状态
			userIdStr, actTypeStr, _ := strings.Cut(item.(string), ":")
			userid, _ := strconv.ParseInt(userIdStr, 10, 64)
			actType, _ := strconv.ParseInt(actTypeStr, 10, 64)

			_, _ = l.svcCtx.UserOptSvcRpcClient.UpdateFavoriteStatus(ctx, &userOptPb.UpdateFavoriteStatusReq{
				VideoId:    videoId,
				UserId:     userid,
				ActionType: actType,
			})
		})
	}, mr.WithWorkers(10))

	return nil
}
