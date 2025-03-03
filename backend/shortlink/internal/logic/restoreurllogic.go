package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestoreUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRestoreUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreUrlLogic {
	return &RestoreUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RestoreUrlLogic) RestoreUrl(in *shortlink.ShortLinkCreateRequest) (*shortlink.ShortLinkCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ShortLinkCreateResponse{}, nil
}
