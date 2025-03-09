package userlogic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"user/pkg/constant"
	"user/pkg/errorcode"

	"user/internal/svc"
	"user/pb/user"

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

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	loginSuccess := l.svcCtx.UserModel.Login(l.ctx, in.Username, in.Password)
	if !loginSuccess {
		return &user.LoginResponse{}, errors.New(errorcode.UserNotExist.Message())
	}
	hasLoginMap, err := l.svcCtx.Redis.Hgetall(constant.USER_LOGIN_KEY + in.GetUsername())
	if err != nil {
		return &user.LoginResponse{}, err
	}

	if len(hasLoginMap) > 0 {
		expireErr := l.svcCtx.Redis.Expire(constant.USER_LOGIN_KEY+in.GetUsername(), constant.USER_LOGIN_EXPIRE_TIME)
		if expireErr != nil {
			logx.Errorf("设置 Redis 过期时间失败: %v", expireErr)
			return &user.LoginResponse{}, expireErr
		}
		token := hasLoginMap["token"]
		return &user.LoginResponse{
			Token: token,
		}, nil
	}

	newToken := uuid.New().String()
	err = l.svcCtx.Redis.Hset(constant.USER_LOGIN_KEY+in.GetUsername(), "token", newToken)
	if err != nil {
		return &user.LoginResponse{}, err
	}

	err = l.svcCtx.Redis.Expire(constant.USER_LOGIN_KEY+in.GetUsername(), constant.USER_LOGIN_EXPIRE_TIME)
	if err != nil {
		logx.Errorf("设置 Redis 过期时间失败: %v", err)
		return &user.LoginResponse{}, err
	}

	return &user.LoginResponse{
		Token: newToken,
	}, nil
}
