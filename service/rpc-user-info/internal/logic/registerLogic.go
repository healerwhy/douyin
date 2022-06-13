package logic

import (
	"context"
	"douyin/common/help/token"
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
	// todo: add your logic here and delete this line
	bytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), 12)
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, nil, &model.User{
		UserName:       in.UserName,
		PasswordDigest: string(bytes),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v insert error", in)
	}
	userId, _ := res.LastInsertId()

	var genToken *token.GenToken
	now := time.Now()
	tokenString, err := genToken.GenToken(now, userId, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "genToken error")
	}
	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, strconv.FormatInt(userId, 10), tokenString, 24*3600)
	if err != nil {
		return nil, errors.Wrapf(err, "set tokenString to redis error")
	}
	return &userInfoPb.RegisterResp{
		UserId: userId,
		Token:  tokenString,
	}, nil
}
