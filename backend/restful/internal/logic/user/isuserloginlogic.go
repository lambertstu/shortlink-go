package user

import (
	"context"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

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

func (l *IsUserLoginLogic) IsUserLogin(req *types.IsUserLoginRequest) (resp *types.NilResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
