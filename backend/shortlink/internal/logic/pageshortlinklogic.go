package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

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
	var list []*shortlink.ShortLinkPageData
	err := l.svcCtx.ShortlinkModel.Pagination(l.ctx, in.GetPage(), in.GetSize(), in.GetOrderTag(), bson.D{}, "full_short_url", list)
	if err != nil {
		return nil, err
	}

	return &shortlink.ShortLinkPageResponse{
		List:  list,
		Total: int32(len(list)),
	}, nil
}
