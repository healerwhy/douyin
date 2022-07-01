package logic

import (
	"context"
	"douyin/service/rpc-user-info/userInfoPb"
	"github.com/jinzhu/copier"
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
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}
	logx.Errorf(" resp: %+v", userInfo)
	/*
		UserId         int64     `db:"user_id"`
		UserName       string    `db:"user_name"`
		FollowCount    int64     `db:"follow_count"`
		FollowerCount  int64     `db:"follower_count"`

		UserId        int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
		UserName      string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
		FollowCount   int64  `protobuf:"varint,3,opt,name=followCount,proto3" json:"followCount,omitempty"`
		FollowerCount int64  `protobuf:"varint,4,opt,name=followerCount,proto3" json:"followerCount,omitempty"`
	*/

	var user userInfoPb.User
	_ = copier.Copy(&user, userInfo)

	return &userInfoPb.UserInfoResp{
		User: &user,
	}, nil
}
