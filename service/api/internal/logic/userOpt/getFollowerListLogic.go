package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-user-info/userInfoPb"
	"douyin/service/rpc-user-operate/useroptservice"
	"github.com/jinzhu/copier"

	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerListLogic) GetFollowerList(req *types.FollowerListReq) (resp *types.FollowerListRes, err error) {

	followersIdMap, err := l.svcCtx.UserOptSvcRpcClient.GetUserFollower(l.ctx, &useroptservice.GetUserFollowerReq{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.FollowerListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get user follower list fail " + err.Error(),
			},
		}, nil
	}

	var followersIdArr []int64
	for k := range followersIdMap.UserFollowerList {
		followersIdArr = append(followersIdArr, k)
	}

	var userList []*types.User // 最终返回的关注者列表

	if followersIdMap != nil {

		followersInfo, err := l.svcCtx.UserInfoRpcClient.AuthsInfo(l.ctx, &userInfoPb.AuthsInfoReq{
			AuthIds: followersIdArr,
		})
		if err != nil {
			return &types.FollowerListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "get AuthsInfo list fail " + err.Error(),
				},
			}, nil
		}

		// 看粉丝有没有关注我
		allFollowersMap, err := l.svcCtx.UserOptSvcRpcClient.GetUserFollow(l.ctx, &useroptservice.GetUserFollowReq{
			UserId:  req.UserId,
			AuthIds: followersIdArr,
		})

		for _, v := range followersInfo.Auths {
			var user types.User
			_ = copier.Copy(&user, v)
			user.IsFollow = allFollowersMap.UserFollowList[v.UserId]

			userList = append(userList, &user)
		}

		return &types.FollowerListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			UserFollowerlist: userList,
		}, nil

	} else { // 没有关注任何人
		return &types.FollowerListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
		}, nil
	}
}
