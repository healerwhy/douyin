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

type GetUserFollowStatusHandler struct {
	svcCtx *svc.ServiceContext
}

func NewGetUserFollowStatusHandler(svcCtx *svc.ServiceContext) *GetUserFollowStatusHandler {
	return &GetUserFollowStatusHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask if return err != nil , asynq will retry
func (l *GetUserFollowStatusHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	vals, err := l.svcCtx.RedisCache.SmembersCtx(ctx, globalkey.FollowSetKey)
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
		for _, followKey := range vals {
			source <- followKey
		}
	}, func(item interface{}) {
		followIdKey := item.(string)
		usersInfoTemp, err := l.svcCtx.RedisCache.EvalShaCtx(ctx, l.svcCtx.ScriptREMTag, []string{followIdKey})

		if err != nil { // 获取赞了这个视频的所有的用户Id
			logx.Errorf("RedisCache.SmembersCtx error -----> %s", err.Error())
			return
		}

		// 切分出视频的Id
		_, followIdStr, _ := strings.Cut(followIdKey, ":")
		followId, _ := strconv.ParseInt(followIdStr, 10, 64)

		var usersInfo []interface{}
		usersInfo = usersInfoTemp.([]interface{})

		mr.ForEach(func(source chan<- interface{}) {
			for _, userId := range usersInfo {
				source <- userId
			}
		}, func(item interface{}) {
			// 切分出用户的Id 和 关注状态 "%d:%d" 0 是未写入数据库 第二个是user_id 第三个是操作 关注与未关注
			members := strings.Split(item.(string), ":")
			userid, _ := strconv.ParseInt(members[0], 10, 64)
			actType, _ := strconv.ParseInt(members[1], 10, 64)

			_, err = l.svcCtx.UserOptSvcRpcClient.UpdateFollowStatus(ctx, &userOptPb.UpdateFollowStatusReq{
				FollowId:   followId,
				UserId:     userid,
				ActionType: actType,
			})
			if err != nil {
				logx.Errorf("UserOptSvcRpcClient.UpdateFollowStatus error -----> %s", err.Error())
				return
			}
		})
	}, mr.WithWorkers(10))

	return nil
}
