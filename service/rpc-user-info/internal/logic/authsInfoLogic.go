package logic

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/rpc-user-info/internal/svc"
	"douyin/service/rpc-user-info/userInfoPb"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthsInfoLogic {
	return &AuthsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthsInfoLogic) AuthsInfo(in *userInfoPb.AuthsInfoReq) (*userInfoPb.AuthsInfoResp, error) {
	whereBuilder := l.svcCtx.UserModel.RowBuilder().Where(squirrel.Eq{"user_id": in.AuthIds})
	/*
		FindAll FindAll 里的降序 需要是user_id 而默认是id 需要改一下源文件
	*/
	auths, err := l.svcCtx.UserModel.FindAll(l.ctx, whereBuilder, "user_id ASC")

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get auths info fail"), "get auths info fail FindAll  err : %v , authIds:%v", err, in.AuthIds)
	}
	/*
	     UserId         int64     `db:"user_id"`
	     UserName       string    `db:"user_name"`
	     PasswordDigest string    `db:"password_digest"`
	     FollowCount    int64     `db:"follow_count"`
	     FollowerCount  int64     `db:"follower_count"`

	   UserId        int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	   UserName      string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	   FollowCount   int64  `protobuf:"varint,3,opt,name=followCount,proto3" json:"followCount,omitempty"`
	   FollowerCount int64  `protobuf:"varint,4,opt,name=followerCount,proto3" json:"followerCount,omitempty"`
	   IsFollowing   bool   `protobuf:"varint,5,opt,name=isFollowing,proto3" json:"isFollowing,omitempty"`

	 SELECT `user_id`,`user_name`,`password_digest`,`follow_count`,`follower_count`,`del_state`,`create_time` FROM `user` WHERE user_id = ('3,4,5,6') AND del_state = 0 ORDER BY user_id ASC;
	*/
	var authsInfo map[int64]*userInfoPb.User
	if len(auths) > 0 {
		authsInfo = make(map[int64]*userInfoPb.User, len(auths))
		for _, v := range auths {
			var authInfo userInfoPb.User
			_ = copier.Copy(&authInfo, v)

			authsInfo[v.UserId] = &authInfo
		}
	}

	return &userInfoPb.AuthsInfoResp{
		Auths: authsInfo,
	}, nil
}
