package shortlink

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageShortLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageShortLinkLogic {
	return &PageShortLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageShortLinkLogic) PageShortLink(req *types.ShortLinkPageRequest) (resp *types.ShortLinkPageResponse, err error) {
	rpcResp, err := l.svcCtx.CoreRpcService.PageShortLink(l.ctx, &shortlinkclient.ShortLinkPageRequest{
		Gid:      req.Gid,
		OrderTag: req.OrderTag,
		Page:     req.Page,
		Size:     req.Size,
	})

	var list []types.ShortLinkPageData
	for _, v := range rpcResp.List {
		list = append(list, types.ShortLinkPageData{
			ShortUri:     v.ShortUri,
			FullShortUrl: v.FullShortUrl,
			OriginUrl:    v.OriginUrl,
			Gid:          v.Gid,
			EnableStatus: v.EnableStatus,
			CreateTime:   v.CreateTime,
			Describe:     v.Describe,
			Favicon:      v.Favicon,
			TotalPv:      v.TotalPv,
			TodayPv:      v.TodayPv,
			TotalUv:      v.TotalUv,
			TodayUv:      v.TodayUv,
			TotalUip:     v.TotalUip,
			TodayUip:     v.TodayUip,
			ClickNum:     v.ClickNum,
		})
	}

	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("获取分组短链接失败")
	}
	return &types.ShortLinkPageResponse{
		List:    list,
		MaxPage: rpcResp.MaxPage,
		Page:    rpcResp.Page,
		Total:   rpcResp.Total,
	}, nil
}
