package userOpt

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/xerr"
	"douyin/service/rpc-user-info/userInfoPb"
	"douyin/service/rpc-user-operate/useroptservice"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"

	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (resp *types.CommentListRes, err error) {
	// 1.获得视频里所有的点赞内容
	// 2.获得点赞用的用户信息，follow情况
	comments, err := l.svcCtx.UserOptSvcRpcClient.GetVideoComment(l.ctx, &useroptservice.GetVideoCommentReq{
		VideoId: req.VideoId,
	})
	if err != nil {
		logx.Errorf("get video comment list fail %s", err.Error())
		return &types.CommentListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get video comment list fail",
			},
		}, nil
	}

	var commentList []*types.Comment // 最终返回的视频列表

	// 当前的用户id
	userId := l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
	if comments != nil {
		// 把视频里的作者的信息找出来 并去重
		var authIds = make([]int64, 0, len(comments.CommentList))
		for _, v := range comments.CommentList {
			var authIdsTemp = make(map[int64]interface{}, len(comments.CommentList))

			if _, ok := authIdsTemp[v.UserId]; !ok {
				authIdsTemp[v.UserId] = nil
				authIds = append(authIds, v.UserId)
			}
		}

		var authsInfo *userInfoPb.AuthsInfoResp
		var userFollowList *useroptservice.GetUserFollowResp
		err = mr.Finish(func() error {
			// 查询所有视频的的作者信息
			authsInfo, err = l.svcCtx.UserInfoRpcClient.AuthsInfo(l.ctx, &userInfoPb.AuthsInfoReq{ // 返回作者信息 按照作者id升序排列
				AuthIds: authIds,
			})
			if err != nil {
				resp = &types.CommentListRes{
					Status: types.Status{
						Code: xerr.ERR,
						Msg:  "get author list fail",
					},
				}
				return err
			}
			return nil
		}, func() error {
			// 查找当前用户对作者的关注状态
			userFollowList, err = l.svcCtx.UserOptSvcRpcClient.GetUserFollow(l.ctx, &useroptservice.GetUserFollowReq{
				UserId:  userId,
				AuthIds: authIds,
			})
			if err != nil {
				resp = &types.CommentListRes{
					Status: types.Status{
						Code: xerr.ERR,
						Msg:  "get user follow list fail",
					},
				}
				return err
			}
			return nil
		})
		if err != nil {
			logx.Errorf("get user follow list fail %s", err.Error())
			return resp, nil
		}

		for _, v := range comments.CommentList {
			var comment types.Comment
			_ = copier.Copy(&comment, v)
			_ = copier.Copy(&comment.User, authsInfo.Auths[v.UserId])
			// 用户对该视频的作者是否关注
			comment.User.IsFollow = userFollowList.UserFollowList[v.UserId]
			comment.CreateTime = v.CreateDate

			commentList = append(commentList, &comment)
		}

		return &types.CommentListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			CommentList: commentList,
		}, nil

	} else { // 没有点赞视频
		return &types.CommentListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
		}, nil
	}

}
