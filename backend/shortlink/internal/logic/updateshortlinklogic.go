package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"shortlink/internal/svc"
	model "shortlink/mongo/shortlink"
	"shortlink/pb/shortlink"
	"shortlink/pkg/constant"
	"time"
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
	err := l.svcCtx.ShortlinkModel.UpdateShortLinkInfo(l.ctx, &model.Shortlink{
		Domain:    constant.CreateShortLinkDefaultDomain,
		Gid:       in.GetGid(),
		OriginUrl: in.GetOriginUrl(),
		Describe:  in.GetDescribe(),
		Favicon:   in.GetFavicon(),
		ClickNum:  int(in.GetClickNum()),
		TotalPv:   int(in.GetTotalPv()),
		TotalUv:   int(in.GetTotalUv()),
		TotalUip:  int(in.GetTotalUip()),
		TodayPv:   int(in.GetTodayPv()),
		TodayUv:   int(in.GetTodayUv()),
		TodayUip:  int(in.GetTodayUip()),
		UpdateAt:  time.Now(),
	})
	if err != nil {
		return &shortlink.ShortLinkUpdateResponse{
			Success: false,
		}, err
	}

	return &shortlink.ShortLinkUpdateResponse{
		Success: true,
	}, nil
}
