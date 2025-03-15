package shortlink

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"

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

func (l *UpdateShortLinkLogic) UpdateShortLink(req *types.ShortLinkUpdateRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.CoreRpcService.UpdateShortLink(l.ctx, &shortlinkclient.ShortLinkUpdateRequest{
		Gid:       req.Gid,
		OriginUrl: req.OriginUrl,
		ShortUri:  req.ShortUri,
		Describe:  req.Describe,
		Favicon:   req.Favicon,
		TodayPv:   req.TodayPv,
		TodayUip:  req.TodayUip,
		TodayUv:   req.TodayUv,
		TotalPv:   req.TotalPv,
		TotalUip:  req.TotalUip,
		TotalUv:   req.TotalUv,
		ClickNum:  req.ClickNum,
	})
	if !rpcResp.Success {
		return nil, errors.New("更新短链接信息失败")
	}
	return nil, nil
}
