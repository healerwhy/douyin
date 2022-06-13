package user

import (
	"context"
	"douyin/common/xerr"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"douyin/service/rpc-user-info/userinfoservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserRegister 对密码使用bcrypt加密
// 生成token，并将token和userId存入redis
// rpc调用Register服务保存用户信息
func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterRes, err error) {
	res, err := l.svcCtx.UserInfoRpcClient.Register(l.ctx, &userinfoservice.RegisterReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return &types.UserRegisterRes{
			Status: types.Status{
				Code: xerr.SECRET_ERROR,
				Msg:  "register failed" + err.Error(),
			},
		}, nil
	}

	return &types.UserRegisterRes{
		Status: types.Status{
			Code: xerr.OK,
		},
		IdWithTokenRes: types.IdWithTokenRes{
			UserId: res.UserId,
			Token:  res.Token,
		},
	}, nil
}
