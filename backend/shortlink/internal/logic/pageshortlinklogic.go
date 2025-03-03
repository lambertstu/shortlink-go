package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageShortLinkLogic {
	return &PageShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PageShortLinkLogic) PageShortLink(in *shortlink.ShortLinkPageRequest) (*shortlink.ShortLinkPageResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ShortLinkPageResponse{}, nil
}
