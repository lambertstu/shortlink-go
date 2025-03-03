package logic

import (
	"context"

	"shortlink/internal/svc"
	"shortlink/pb/shortlink"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupShortLinkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupShortLinkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupShortLinkCountLogic {
	return &ListGroupShortLinkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListGroupShortLinkCountLogic) ListGroupShortLinkCount(in *shortlink.ListGroupShortLinkCountRequest) (*shortlink.ListGroupShortLinkCountResponse, error) {
	// todo: add your logic here and delete this line

	return &shortlink.ListGroupShortLinkCountResponse{}, nil
}
