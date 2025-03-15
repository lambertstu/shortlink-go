package user

import (
	"context"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	rpcResp, rpcErr := l.svcCtx.UserRpcService.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if rpcErr != nil {
		return nil, rpcErr
	}
	return &types.LoginResponse{
		Token: rpcResp.Token,
	}, nil
}
