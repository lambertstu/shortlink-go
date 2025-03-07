package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"shortlink/internal/svc"
	model "shortlink/mongo/shortlink"
	"shortlink/pb/shortlink"
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
		Domain:    in.GetDomain(),
		OriginUrl: in.GetOriginUrl(),
		Gid:       in.GetGid(),
		Describe:  in.GetDescribe(),
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
