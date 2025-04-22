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

	// 创建成功后，查询短链接信息以获取ShortUri
	// 因为RPC返回中可能没有直接包含ShortUri
	shortlinkInfo, qErr := l.svcCtx.ShortLinkModel.FindOneByOriginUrl(l.ctx, req.OriginUrl)
	if qErr == nil && shortlinkInfo != nil {
		// 获取到短链接URI，添加到布隆过滤器
		shortUri := shortlinkInfo.ShortUri
		l.Logger.Infof("短链接创建成功: %s -> %s", req.OriginUrl, shortUri)

		// 将短链接添加到布隆过滤器中
		if err := l.svcCtx.BloomFilter.Add([]byte(shortUri)); err != nil {
			l.Logger.Errorf("将短链接添加到布隆过滤器失败: %v", err)
		} else {
			l.Logger.Infof("已将短链接 %s 添加到布隆过滤器", shortUri)
		}
	} 
	}

	return &types.ShortLinkCreateResponse{
		OriginUrl: rpcResp.OriginUrl,
	}, nil
}
