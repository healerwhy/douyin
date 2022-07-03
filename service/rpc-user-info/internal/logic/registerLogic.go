package logic

import (
	"context"
	"douyin/common/help/token"
	"douyin/common/xerr"
	"douyin/service/rpc-user-info/model"
	"douyin/service/rpc-user-info/userInfoPb"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"douyin/service/rpc-user-info/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register -----------------------user-----------------------
func (l *RegisterLogic) Register(in *userInfoPb.RegisterReq) (*userInfoPb.RegisterResp, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), 12)
	if err != nil {
		logx.Errorf("generate password failed, err:%s", err.Error())
		return nil, err
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, nil, &model.User{
		UserName:       in.UserName,
		PasswordDigest: string(bytes),
	})
	if err != nil {
		logx.Errorf("insert user failed, err: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert user failed, user_name: %s", in.UserName)
	}
	userId, _ := res.LastInsertId()

	var genToken *token.GenToken
	now := time.Now()
	tokenString, err := genToken.GenToken(now, userId, nil)
	if err != nil {
		logx.Errorf("gen token error: %s", err.Error())
		return nil, errors.Wrapf(err, "genToken error")
	}

	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, "token:"+strconv.FormatInt(userId, 10), tokenString, token.AccessExpire)
	if err != nil {
		logx.Errorf("set token to redis error: %s", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "genToken error")
	}

	return &userInfoPb.RegisterResp{
		UserId: userId,
		Token:  tokenString,
	}, nil
}
