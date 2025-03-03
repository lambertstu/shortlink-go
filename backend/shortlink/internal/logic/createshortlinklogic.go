package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShortLinkLogic {
	return &CreateShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateShortLinkLogic) CreateShortLink(in *shortlink.ShortLinkCreateRequest) (*shortlink.ShortLinkCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ShortLinkCreateResponse{}, nil
}
