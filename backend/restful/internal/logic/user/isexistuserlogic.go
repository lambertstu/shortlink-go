package user

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExistUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsExistUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistUserLogic {
	return &IsExistUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsExistUserLogic) IsExistUser(req *types.IsExistUserRequest) (resp *types.IsExistUserResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
