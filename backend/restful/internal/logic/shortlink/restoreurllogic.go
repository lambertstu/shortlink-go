package shortlink

import (
	"context"
	"errors"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestoreUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestoreUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreUrlLogic {
	return &RestoreUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestoreUrlLogic) RestoreUrl(req *types.ShortLinkRestoreRequest) (resp *types.ShortLinkRestoreResponse, err error) {
	shortlinkInfo, err := l.svcCtx.ShortLinkModel.FindOneByShortUrl(l.ctx, req.ShortUri)
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("短链接查询失败")
	}

	return &types.ShortLinkRestoreResponse{
		OriginUrl: shortlinkInfo.OriginUrl,
	}, nil
}
