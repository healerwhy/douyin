package user

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"douyin/service/rpc-user-info/userinfoservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserLogin 用户登陆
// 通过username获得用户密码，然后比对密码
// token 先查redis 是否存在，如果存在，则直接返回，如果不存在，则生成token，并存入redis
// 并返回userId，token
func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginRes, err error) {
	res, err := l.svcCtx.UserInfoRpcClient.Login(l.ctx, &userinfoservice.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("login failed: %v", err.Error())
		return &types.UserLoginRes{
			Status: types.Status{
				Code: xerr.ERR,
				Msg:  "login failed",
			},
		}, nil
	}

	return &types.UserLoginRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		IdWithTokenRes: types.IdWithTokenRes{
			UserId: res.UserId,
			Token:  res.Token,
		},
	}, nil
}
