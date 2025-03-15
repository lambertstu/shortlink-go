package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

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
	rpcResp, _ := l.svcCtx.UserRpcService.IsUserLogin(l.ctx, &user.IsUserLoginRequest{
		Token:    req.Token,
		Username: req.Username,
	})
	if !rpcResp.Success {
		return nil, errors.New("查询用户登录情况失败")
	}
	return nil, nil
}
