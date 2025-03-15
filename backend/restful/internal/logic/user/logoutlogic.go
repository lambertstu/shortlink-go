package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.UserRpcService.Logout(l.ctx, &user.LogoutRequest{
		Token:    req.Token,
		Username: req.Username,
	})
	if !rpcResp.Success {
		return nil, errors.New("用户注销失败")
	}
	return nil, nil
}
