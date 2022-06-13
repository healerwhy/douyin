package logic

import (
	"context"
	myToken "douyin/common/help/token"
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
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.UserName)
	if err != nil {
		return nil, errors.Wrap(err, "find user failed")
	}

	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(in.Password))
	if err != nil {
		return nil, errors.Wrapf(err, "password not match")
	}

	// 通过userId查 redis
	token, err := l.svcCtx.RedisCache.GetCtx(l.ctx, strconv.FormatInt(user.UserId, 10))
	if err != nil {
		return nil, errors.Wrap(err, "get token from redis failed")
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
	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, strconv.FormatInt(user.UserId, 10), token, 60*60*24)
	if err != nil {
		return nil, errors.Wrapf(err, "set token to redis error")
	}

	return &userInfoPb.LoginResp{
		UserId: user.UserId,
		Token:  token,
	}, nil
}
