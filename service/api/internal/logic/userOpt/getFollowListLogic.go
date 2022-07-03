package userOpt

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"douyin/service/rpc-user-info/userInfoPb"
	"douyin/service/rpc-user-operate/useroptservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowListLogic) GetFollowList(req *types.FollowListReq) (resp *types.FollowListRes, err error) {

	followsIdMap, err := l.svcCtx.UserOptSvcRpcClient.GetUserFollow(l.ctx, &useroptservice.GetUserFollowReq{
		UserId: req.UserId,
	})
	if err != nil {
		logx.Errorf("get user follow list fail %s", err.Error())
		return &types.FollowListRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "get user follow list fail ",
			},
		}, nil
	}

	var followsIdArr []int64
	for k := range followsIdMap.UserFollowList {
		followsIdArr = append(followsIdArr, k)
	}

	var userList []*types.User // 最终返回的关注者列表

	if followsIdMap != nil {

		followsInfo, err := l.svcCtx.UserInfoRpcClient.AuthsInfo(l.ctx, &userInfoPb.AuthsInfoReq{
			AuthIds: followsIdArr,
		})
		if err != nil {
			logx.Errorf("get user follow list fail %s", err.Error())
			return &types.FollowListRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  "get AuthsInfo list fail " + err.Error(),
				},
			}, nil
		}

		for _, v := range followsInfo.Auths {
			var user types.User
			_ = copier.Copy(&user, v)
			user.IsFollow = true

			userList = append(userList, &user)
		}

		return &types.FollowListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
			UserFollowlist: userList,
		}, nil

	} else { // 没有关注任何人
		return &types.FollowListRes{
			Status: types.Status{
				Code: xerr.OK,
			},
		}, nil
	}
}
