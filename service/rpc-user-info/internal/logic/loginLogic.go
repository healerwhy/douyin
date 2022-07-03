package logic

import (
	"context"
	myToken "douyin/common/help/token"
	"douyin/common/xerr"
	"douyin/service/rpc-user-info/internal/svc"
	"douyin/service/rpc-user-info/userInfoPb"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 登录
// 然后通过username获得用户密码，然后比对密码
// 通过userId查 redis
// 如果存在，则直接返回，如果不存在，则生成token，并存入redis
func (l *LoginLogic) Login(in *userInfoPb.LoginReq) (*userInfoPb.LoginResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.UserName)
	if err != nil {
		logx.Errorf("find user failed, err: %s", err.Error())
		return nil, errors.Wrap(err, "find user failed")
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(in.Password))
	if err != nil {
		logx.Errorf("password not match, err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "password not match")
	}

	// 通过userId查 redis 是否有此token
	token, err := l.svcCtx.RedisCache.GetCtx(l.ctx, "token:"+strconv.FormatInt(user.UserId, 10))
	if err != nil {
		logx.Errorf("get token from redis failed, err: %s", err.Error())
		return nil, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get token from redis failed")
	}
	// 如果存在，则直接返回
	if token != "" {
		return &userInfoPb.LoginResp{
			UserId: user.UserId,
			Token:  token,
		}, nil
	}

	//如果不存在，则生成token，并存入redis
	var genToken myToken.GenToken
	now := time.Now()
	token, err = genToken.GenToken(now, user.UserId, nil)
	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, "token:"+strconv.FormatInt(user.UserId, 10), token, myToken.AccessExpire)
	if err != nil {
		logx.Errorf("set token to redis failed, err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "set token to redis error")
	}

	return &userInfoPb.LoginResp{
		UserId: user.UserId,
		Token:  token,
	}, nil
}
