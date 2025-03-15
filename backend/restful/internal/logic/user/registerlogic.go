package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.UserRpcService.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		RealName: req.RealName,
		Phone:    req.Phone,
	})
	if !rpcResp.Success {
		return nil, errors.New("用户注册失败,用户名已使用")
	}
	return nil, nil
}
