package logic

import (
	"context"
	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/model"
	"douyin/service/rpc-user-operate/userOptPb"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFavoriteStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFavoriteStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFavoriteStatusLogic {
	return &UpdateFavoriteStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// export logic "INSERT INTO user_favorite_list(user_id,video_id,is_favorite) VALUES(1,6,?) ON DUPLICATE KEY UPDATE is_favorite=?"
func (l *UpdateFavoriteStatusLogic) UpdateFavoriteStatus(in *userOptPb.UpdateFavoriteStatusReq) (*userOptPb.UpdateFavoriteStatusResp, error) {
	tmp := []string{"user_id", "video_id", "is_favorite"}
	field := strings.Join(tmp, ",")
	err := l.svcCtx.UserFavoriteModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.UserFavoriteModel.InsertOrUpdate(l.ctx, session, field, "is_favorite", in.UserId, in.VideoId, in.ActionType)
		if err != nil {
			logx.Errorf("UpdateFavoriteStatusLogic------->InsertOrUpdate err : %v", err.Error())
			return err
		}

		// 消息中传来的 in.action是 0 1 写入user favorite_count 就需要变成 -1 1
		action := l.getActionType(in.ActionType)
		_, err = l.svcCtx.VideoModel.UpdateStatus(l.ctx, session, "favorite_count", "id", action, in.VideoId)
		if err != nil {
			logx.Errorf("UpdateFavoriteStatusLogic------->UpdateStatus err : %v", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		logx.Error("UpdateFavoriteStatusLogic-------> trans fail")
		return &userOptPb.UpdateFavoriteStatusResp{}, err
	}

	return &userOptPb.UpdateFavoriteStatusResp{}, nil
}
func (l *UpdateFavoriteStatusLogic) getActionType(actionType int64) int64 {

	switch actionType { // 方便扩展
	case model.ActionADD:
		return 1
	case model.ActionCancel:
		return -1
	default:
		return -99
	}
}
