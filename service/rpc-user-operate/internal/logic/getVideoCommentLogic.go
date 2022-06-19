package logic

import (
	"context"
	"douyin/common/help/cos"
	"douyin/service/rpc-user-operate/internal/svc"
	"douyin/service/rpc-user-operate/userOptPb"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentLogic {
	return &GetVideoCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoCommentLogic) GetVideoComment(in *userOptPb.GetVideoCommentReq) (*userOptPb.GetVideoCommentReqResp, error) {

	var downloadHelper cos.DownloadComment
	key := fmt.Sprintf("video_id_%d/", in.VideoId)
	comments, err := downloadHelper.DownloadComment(l.ctx, key, l.svcCtx.Config.COSConf.CommentBucket, l.svcCtx.Config.COSConf.SecretId, l.svcCtx.Config.COSConf.SecretKey)

	var commentList []*userOptPb.Comment
	for _, v := range comments {
		var comment userOptPb.Comment
	_:
		copier.Copy(&comment, v)
		comment.Content = v.Content
		commentList = append(commentList, &comment)
	}

	if err != nil {
		return nil, err
	}

	return &userOptPb.GetVideoCommentReqResp{
		CommentList: commentList,
	}, nil
}
