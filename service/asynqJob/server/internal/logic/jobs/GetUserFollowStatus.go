package jobs

import (
	"context"
	"douyin/service/asynqJob/server/internal/svc"
	"fmt"
	"github.com/hibiken/asynq"
)

// GetUserFollowStatusHandler   shcedule billing to home business
type GetUserFollowStatusHandler struct {
	svcCtx *svc.ServiceContext
}

func NewGetUserFollowStatusHandler(svcCtx *svc.ServiceContext) *GetUserFollowStatusHandler {
	return &GetUserFollowStatusHandler{
		svcCtx: svcCtx,
	}
}

//  every one minute exec : if return err != nil , asynq will retry
func (l *GetUserFollowStatusHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	fmt.Printf("shcedule server demo -----> every one minute exec \n")

	return nil
}
