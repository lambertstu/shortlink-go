package shortlink

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"

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

func (l *DeleteShortLinkLogic) DeleteShortLink(req *types.ShortLinkDeleteRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.CoreRpcService.DeleteShortLink(l.ctx, &shortlinkclient.ShortLinkDeleteRequest{
		OriginUrl: req.OriginUrl,
		ShortUri:  req.ShortUri,
	})
	if !rpcResp.Success {
		return nil, errors.New("删除短链接失败")
	}
	return nil, nil
}
