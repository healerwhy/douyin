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

	vals, err := l.svcCtx.RedisCache.SmembersCtx(ctx, globalkey.FavoriteSetKey)
	if err != nil {
		logx.Errorf("RedisCache.SmembersCtx error -----> %s", err.Error())
		return err
	}
	if len(vals) == 0 {
		logx.Infof("every 10s exec But not exist data in redis cache")
		return nil
	}

	// 持久化数据
	mr.ForEach(func(source chan<- interface{}) {
		for _, videoIdKey := range vals {
			source <- videoIdKey
		}
	}, func(item interface{}) {
		videoIdKey := item.(string)

		usersInfoTemp, err := l.svcCtx.RedisCache.EvalShaCtx(ctx, l.svcCtx.ScriptREMTag, []string{videoIdKey})

		if err != nil { // 获取赞了这个视频的所有的用户Id
			logx.Errorf("RedisCache.SmembersCtx error -----> %s", err.Error())
			return
		}

		// 切分出视频的Id
		_, videoIdStr, _ := strings.Cut(videoIdKey, ":")
		videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

		var usersInfo []interface{}
		usersInfo = usersInfoTemp.([]interface{})

		mr.ForEach(func(source chan<- interface{}) {
			for _, userId := range usersInfo {
				source <- userId
			}
		}, func(item interface{}) {
			// 切分出用户的Id 和 点赞状态 "%d:%d" 0 是未写入数据库 第二个是user_id 第三个是操作 点赞与未点赞
			members := strings.Split(item.(string), ":")
			userid, _ := strconv.ParseInt(members[0], 10, 64)
			actType, _ := strconv.ParseInt(members[1], 10, 64)
			_, _ = l.svcCtx.UserOptSvcRpcClient.UpdateFavoriteStatus(ctx, &userOptPb.UpdateFavoriteStatusReq{
				VideoId:    videoId,
				UserId:     userid,
				ActionType: actType,
			})
		})
	}, mr.WithWorkers(10))

	return nil
}
