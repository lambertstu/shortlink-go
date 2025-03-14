package shortlink

import (
	"context"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShortLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShortLinkLogic {
	return &DeleteShortLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteShortLinkLogic) DeleteShortLink(req *types.ShortLinkDeleteRequest) (resp *types.ShortLinkDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
