package logic

import (
	"context"
	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExistUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistUserLogic {
	return &IsExistUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsExistUserLogic) IsExistUser(in *user.IsExistUserRequest) (*user.IsExistUserResponse, error) {
	exists, err := l.svcCtx.BloomFilter.Exists([]byte(in.GetUsername()))
	if err != nil || !exists {
		return &user.IsExistUserResponse{
			Success: false,
		}, err
	}

	return &user.IsExistUserResponse{
		Success: true,
	}, nil
}
