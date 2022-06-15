package userOpt

import (
	"context"

	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentOptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentOptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentOptLogic {
	return &CommentOptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentOptLogic) CommentOpt(req *types.CommentOptReq) (resp *types.CommentOptRes, err error) {
	// todo: add your logic here and delete this line

	return
}
