package jobs

import (
	"context"
	"douyin/service/asynqJob/server/internal/svc"
	"fmt"
	"github.com/hibiken/asynq"
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

//  every one minute exec : if return err != nil , asynq will retry
func (l *GetUserFavoriteStatusHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	fmt.Printf("shcedule server demo -----> NewGetUserFavoriteStatusHandler \n")

	return nil
}
