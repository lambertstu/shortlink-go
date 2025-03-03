package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateShortLinkLogic {
	return &BatchCreateShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchCreateShortLinkLogic) BatchCreateShortLink(in *shortlink.ShortLinkBatchCreateRequest) (*shortlink.ShortLinkCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ShortLinkCreateResponse{}, nil
}
