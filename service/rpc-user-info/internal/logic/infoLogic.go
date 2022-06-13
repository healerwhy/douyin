package logic

import (
	"context"
	"douyin/service/rpc-user-info/userInfoPb"
	"github.com/pkg/errors"

	"douyin/service/rpc-user-info/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *userInfoPb.UserInfoReq) (*userInfoPb.UserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}

	// 这里的isfollowing暂时定成false 后面应该根据 关系表进行再次查询
	return &userInfoPb.UserInfoResp{
		User: &userInfoPb.User{
			UserId:        user.UserId,
			UserName:      user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollowing:   false,
		},
	}, nil
}
