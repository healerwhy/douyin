package logic

import (
	"context"
	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
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
	_, err := l.svcCtx.UserFavorite.InsertOrUpdate(l.ctx, nil, field, "is_favorite", in.UserId, in.VideoId, in.ActionType)
	if err != nil {
		logx.Errorf("UpdateFavoriteStatusLogic------->UpdateFavoriteStatus err : %v\n", err)
		return nil, nil
	}
	return &userOptPb.UpdateFavoriteStatusResp{}, nil
}
