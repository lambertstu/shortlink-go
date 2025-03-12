package user

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserLoginLogic {
	return &IsUserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsUserLoginLogic) IsUserLogin(req *types.IsUserLoginRequest) (resp *types.IsUserLoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
