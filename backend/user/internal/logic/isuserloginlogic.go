package logic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserLoginLogic {
	return &IsUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsUserLoginLogic) IsUserLogin(in *user.IsUserLoginRequest) (*user.IsUserLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.IsUserLoginResponse{}, nil
}
