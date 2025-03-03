package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShortLinkLogic {
	return &UpdateShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateShortLinkLogic) UpdateShortLink(in *shortlink.ShortLinkUpdateRequest) (*shortlink.ShortLinkUpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ShortLinkUpdateResponse{}, nil
}
