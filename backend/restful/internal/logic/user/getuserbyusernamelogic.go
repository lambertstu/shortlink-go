package user

import (
	"context"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	userInfo, err := l.svcCtx.UserRpcService.GetUserByUsername(l.ctx, &user.GetUserRequest{
		Username: req.Username,
	})

	if err != nil {
		return nil, err
	}

	return &types.GetUserResponse{
		Username: userInfo.Username,
		RealName: userInfo.RealName,
		Phone:    userInfo.Phone,
		Email:    userInfo.Email,
	}, nil
}
