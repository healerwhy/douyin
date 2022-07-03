package logic

import (
	"context"
	"douyin/service/rpc-user-operate/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"

	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFollowStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFollowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFollowStatusLogic {
	return &UpdateFollowStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFollowStatusLogic) UpdateFollowStatus(in *userOptPb.UpdateFollowStatusReq) (*userOptPb.UpdateFollowStatusResp, error) {

	tmp := []string{"user_id", "follow_id", "is_follow"}
	field := strings.Join(tmp, ",")
	err := l.svcCtx.UserFollowModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.UserFollowModel.InsertOrUpdate(l.ctx, session, field, "is_follow", in.UserId, in.FollowId, in.ActionType)
		if err != nil {
			logx.Errorf("UpdateFollowStatus------->InsertOrUpdate err : %s", err.Error())
			return err
		}

		// 消息中传来的 in.action是 0 1 写入user follow_count 就需要变成 -1 / +1
		action := l.getActionType(in.ActionType)
		_, err = l.svcCtx.UserModel.UpdateStatus(l.ctx, session, "follow_count", "user_id", action, in.UserId)
		if err != nil {
			logx.Errorf("UpdateFollowStatus------->UpdateStatus err : %s", err.Error())
			return err
		}

		// 更新被关注者的粉丝数
		_, err = l.svcCtx.UserModel.UpdateStatus(l.ctx, session, "follower_count", "user_id", action, in.FollowId)
		if err != nil {
			logx.Errorf("UpdateFollowStatus------->Update follower Status err : %v", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		logx.Error("UpdateFollowStatus-------> trans fail")
		return &userOptPb.UpdateFollowStatusResp{}, err
	}

	return &userOptPb.UpdateFollowStatusResp{}, nil
}

func (l *UpdateFollowStatusLogic) getActionType(actionType int64) int64 {

	switch actionType { // 方便扩展
	case model.ActionADD:
		return 1
	case model.ActionCancel:
		return -1
	default:
		return -99
	}
}
