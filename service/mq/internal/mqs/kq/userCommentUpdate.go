package kq

import (
	"context"
	"douyin/common/help/cos"
	"douyin/common/messageTypes"
	"douyin/service/mq/internal/svc"
	"douyin/service/rpc-user-operate/model"
	"douyin/service/rpc-user-operate/useroptservice"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
	Listening to the payment flow status change notification message queue
*/
type UserCommentOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	CommentOpt cos.UpdateComment
}

func NewUserCommentUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserCommentOpt {
	return &UserCommentOpt{
		ctx:        ctx,
		svcCtx:     svcCtx,
		CommentOpt: cos.UpdateComment{},
	}
}

func (l *UserCommentOpt) Consume(_, val string) error {
	var message messageTypes.UserCommentOptMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserCommentOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserCommentOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

// 处理逻辑
func (l *UserCommentOpt) execService(message messageTypes.UserCommentOptMessage) error {

	actionType, err := l.getActionType(message.ActionType, message)
	if actionType == -99 || err != nil {
		return errors.Wrap(err, "UserCommentOptMessage->execService getActionType err")
	}

	// 调用rpc 更新user_comment表
	_, err = l.svcCtx.UserOptSvcRpcClient.UpdateCommentStatus(l.ctx, &useroptservice.UpdateCommentStatusReq{
		VideoId:    message.VideoId,
		UserId:     message.UserId,
		CommentId:  message.CommentId,
		ActionType: actionType,
	})

	logx.Error("UserCommentOptMessage->execService xxxxxxxxxxx")

	if err != nil {
		logx.Errorf("UserCommentOptMessage->execService  err : %v , val : %s , message:%+v", err, message)
		return err
	}

	return nil
}

func (l *UserCommentOpt) getActionType(actionType int64, message messageTypes.UserCommentOptMessage) (int64, error) {

	_ = copier.Copy(&l.CommentOpt, &message)

	switch actionType { // 方便扩展
	case model.ActionADD:
		// 新增评论
		_, err := l.CommentOpt.UploadComment(l.ctx, l.svcCtx.Config.COSConf.CommentBucket,
			l.svcCtx.Config.COSConf.SecretId, l.svcCtx.Config.COSConf.SecretKey)

		if err != nil {
			logx.Errorf("UserCommentOptMessage->getActionType  err : %v , val : %s , message:%+v", err, message)
			return 0, err
		}

		return model.ActionADD, nil

	case model.ActionCancel:
		// 取消评论
		_, err := l.CommentOpt.DeleteComment(l.ctx, l.svcCtx.Config.COSConf.CommentBucket,
			l.svcCtx.Config.COSConf.SecretId, l.svcCtx.Config.COSConf.SecretKey)

		if err != nil {
			logx.Errorf("UserCommentOptMessage->getActionType  err : %v , val : %s , message:%+v", err, message)
			return 0, err
		}

		return model.ActionCancel, nil

	default:
		return -99, errors.New("UserCommentOptMessage->execService ActionType err")
	}
}
