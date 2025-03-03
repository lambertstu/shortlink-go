package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShortLinkByLockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateShortLinkByLockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShortLinkByLockLogic {
	return &CreateShortLinkByLockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateShortLinkByLockLogic) CreateShortLinkByLock(in *shortlink.ShortLinkCreateRequest) (*shortlink.ShortLinkCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ShortLinkCreateResponse{}, nil
}
