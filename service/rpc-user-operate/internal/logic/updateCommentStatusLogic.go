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

type UpdateCommentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentStatusLogic {
	return &UpdateCommentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentStatusLogic) UpdateCommentStatus(in *userOptPb.UpdateCommentStatusReq) (*userOptPb.UpdateCommentStatusResp, error) {
	tmp := []string{"video_id", "comment_id", "user_id", "del_state"}
	field := strings.Join(tmp, ",")
	var action int64
	if in.ActionType == model.ActionADD { // 1是添加评论 0是删除评论
		action = 0 // del_state 0 是存在 1是已删除
	} else if in.ActionType == model.ActionCancel {
		action = 1
	}

	err := l.svcCtx.UserCommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {

		// 这里有点不一样 要是model文件里的内容变了 这里也要变
		// InsertOrUpdate(ctx context.Context, session sqlx.Session, field string, setStatus string, videoId, objId, userId, opt int64)
		_, err := l.svcCtx.UserCommentModel.InsertOrUpdate(l.ctx, session, field, "del_state", in.VideoId, in.CommentId, in.UserId, action)
		if err != nil {
			logx.Errorf("UpdateCommentStatus------->InsertOrUpdate err : %s", err.Error())
			return err
		}

		// 消息中传来的 in.action是 0 1 写入video comment_count 就需要变成 -1 / +1
		actionType := l.getActionType(in.ActionType)
		_, err = l.svcCtx.VideoModel.UpdateStatus(l.ctx, session, "comment_count", "id", actionType, in.VideoId)
		if err != nil {
			logx.Errorf("UpdateCommentStatus------->UpdateStatus err : %s", err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logx.Error("UpdateCommentStatus-------> trans fail")
		return &userOptPb.UpdateCommentStatusResp{}, err
	}

	return &userOptPb.UpdateCommentStatusResp{}, nil
}

func (l *UpdateCommentStatusLogic) getActionType(actionType int64) int64 {

	switch actionType { // 方便扩展
	case model.ActionADD:
		return 1
	case model.ActionCancel:
		return -1
	default:
		return -99
	}
}
