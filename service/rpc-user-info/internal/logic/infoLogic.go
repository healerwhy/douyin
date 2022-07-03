package logic

import (
	"context"
	"douyin/common/xerr"
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
		logx.Errorf("get user info failed: %v", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "req: %+v", in)
	}

	var user userInfoPb.User
	_ = copier.Copy(&user, userInfo)

	return &userInfoPb.UserInfoResp{
		User: &user,
	}, nil
}
