package userlogic

import (
	"context"
	"errors"
	"user/pkg/constant"
	"user/pkg/errorcode"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutRequest) (*user.LogoutResponse, error) {
	hasLoginMap, _ := l.svcCtx.Redis.Hgetall(constant.USER_LOGIN_KEY + in.GetUsername())
	if len(hasLoginMap) == 0 || hasLoginMap["token"] != in.Token {
		return &user.LogoutResponse{
			Success: false,
		}, errors.New(errorcode.IdempotentTokenDeleteError.Message())
	}

	_, err := l.svcCtx.Redis.Del(constant.USER_LOGIN_KEY + in.GetUsername())
	if err != nil {
		return &user.LogoutResponse{
			Success: false,
		}, err
	}
	return &user.LogoutResponse{
		Success: true,
	}, nil
}
