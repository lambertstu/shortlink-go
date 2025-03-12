package user

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

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

func (l *UpsertUserInfoLogic) UpsertUserInfo(req *types.UpsertUserInfoRequest) (resp *types.UpsertUserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
