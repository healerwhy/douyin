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

	logx.Infof("NewGetUserFavoriteStatusHandler server -----> every 20s exec ")
	vals, err := l.svcCtx.RedisCache.SmembersCtx(ctx, globalkey.FavoriteSetKey)
	if err != nil {
		logx.Errorf("RedisCache.SmembersCtx error -----> %v", err)
		return err
	}
	if len(vals) == 0 {
		logx.Infof("RedisCache.SmembersCtx no data")
		return nil
	}

	// 持久化数据
	mr.ForEach(func(source chan<- interface{}) {
		for _, videoIdKey := range vals {
			source <- videoIdKey
		}
	}, func(item interface{}) {
		videoIdKey := item.(string)

		// 从拉下来的东西都删掉users, err :=
		usersInfoTemp, err := l.svcCtx.RedisCache.EvalShaCtx(ctx, l.svcCtx.ScriptREMTag, []string{videoIdKey})

		logx.Infof("RedisCache.EvalShaCtx error -----> %v %v", err, usersInfoTemp)

		if err != nil { // 获取赞了这个视频的所有的用户Id
			logx.Errorf("RedisCache.SmembersCtx error -----> %v", err)
			return
		}

		// 切分出视频的Id
		_, videoIdStr, _ := strings.Cut(videoIdKey, ":")
		videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

		logx.Infof("RedisCache.EvalShaCtx usersInfo +++++ %v", usersInfoTemp)

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
