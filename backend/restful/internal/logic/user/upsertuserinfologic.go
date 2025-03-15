package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpsertUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertUserInfoLogic {
	return &UpsertUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpsertUserInfoLogic) UpsertUserInfo(req *types.UpsertUserInfoRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.UserRpcService.UpsertUserInfo(l.ctx, &user.UpsertUserInfoRequest{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		RealName: req.RealName,
		Email:    req.Email,
	})
	if !rpcResp.Success {
		return nil, errors.New("用户信息更新失败")
	}
	return nil, nil
}
