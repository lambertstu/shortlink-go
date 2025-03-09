package userlogic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUsernameLogic {
	return &GetUserByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUsernameLogic) GetUserByUsername(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	userInfo, err := l.svcCtx.UserModel.GetUserByUserName(l.ctx, in.Username)
	if err != nil {
		return &user.GetUserResponse{}, err
	}
	return &user.GetUserResponse{
		Username: userInfo.Username,
		RealName: userInfo.RealName,
		Phone:    userInfo.Phone,
		Email:    userInfo.Mail,
	}, nil
}
