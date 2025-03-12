package shortlink

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageShortLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageShortLinkLogic {
	return &PageShortLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageShortLinkLogic) PageShortLink(req *types.ShortLinkPageRequest) (resp *types.ShortLinkPageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
