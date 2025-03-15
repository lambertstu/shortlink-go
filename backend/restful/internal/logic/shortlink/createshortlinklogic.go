package shortlink

import (
	"context"
	"errors"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShortLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShortLinkLogic {
	return &CreateShortLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateShortLinkLogic) CreateShortLink(req *types.ShortLinkCreateRequest) (resp *types.ShortLinkCreateResponse, err error) {
	rpcResp, err := l.svcCtx.CoreRpcService.CreateShortLink(l.ctx, &shortlinkclient.ShortLinkCreateRequest{
		OriginUrl: req.OriginUrl,
		Gid:       req.Gid,
		Describe:  req.Describe,
	})

	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	if !rpcResp.Success {
		return nil, errors.New("短链接生成失败")
	}

	return &types.ShortLinkCreateResponse{
		OriginUrl: rpcResp.OriginUrl,
	}, nil
}
