package shortlink

import (
	"context"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShortLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShortLinkLogic {
	return &UpdateShortLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShortLinkLogic) UpdateShortLink(req *types.ShortLinkUpdateRequest) (resp *types.ShortLinkUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
