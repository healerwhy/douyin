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
		logx.Errorf("get auths info fail FindAll  err : %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get auths info fail FindAll  err : %v , authIds:%v", err, in.AuthIds)
	}

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
