package tasks

import (
	"context"
	"douyin/service/asynqTask/server/internal/svc"
	"fmt"
	"github.com/hibiken/asynq"
)

// GetUserCommentHandler   shcedule billing to home business
type GetUserCommentHandler struct {
	svcCtx *svc.ServiceContext
}

func NewGetUserCommentHandler(svcCtx *svc.ServiceContext) *GetUserCommentHandler {
	return &GetUserCommentHandler{
		svcCtx: svcCtx,
	}
}

//  every one minute exec : if return err != nil , asynq will retry
func (l *GetUserCommentHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	fmt.Printf("shcedule server demo -----> every one minute exec \n")

	return nil
}
