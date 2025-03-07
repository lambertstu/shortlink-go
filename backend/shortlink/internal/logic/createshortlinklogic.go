package logic

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"shortlink/internal/svc"
	model "shortlink/mongo/shortlink"
	"shortlink/pb/shortlink"
	"shortlink/pkg/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

const createShortLinkDefaultDomain = "zero"

type CreateShortLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateShortLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShortLinkLogic {
	return &CreateShortLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateShortLinkLogic) CreateShortLink(in *shortlink.ShortLinkCreateRequest) (*shortlink.ShortLinkCreateResponse, error) {
	// TODO 缓存
	shortLinkSuffix, err := l.generateSuffix(in)
	if err != nil {
		return nil, err
	}

	fullShortUrl := createShortLinkDefaultDomain + "/" + shortLinkSuffix
	err = l.svcCtx.ShortlinkModel.InsertOneShortlink(l.ctx, &model.Shortlink{
		Domain:       createShortLinkDefaultDomain,
		ShortUri:     shortLinkSuffix,
		FullShortUrl: fullShortUrl,
		OriginUrl:    in.OriginUrl,
		Gid:          in.Gid,
		Describe:     in.Describe,
	})

	if errors.Is(err, model.ErrShortUriExist) {
		bloomErr := l.svcCtx.BloomFilter.Add([]byte(shortLinkSuffix))
		if err != nil {
			return nil, bloomErr
		}

		// TODO 日志
		return &shortlink.ShortLinkCreateResponse{
			Domain:    createShortLinkDefaultDomain,
			OriginUrl: in.OriginUrl,
			Success:   false,
		}, nil
	}

	return &shortlink.ShortLinkCreateResponse{
		Domain:    createShortLinkDefaultDomain,
		OriginUrl: in.OriginUrl,
		Success:   true,
	}, nil
}

func (l *CreateShortLinkLogic) generateSuffix(in *shortlink.ShortLinkCreateRequest) (string, error) {
	var (
		shortUri            string
		customGenerateCount = 0
	)

	for {
		if customGenerateCount > 10 {
			return "", errors.New("短链接频繁生成，请稍后再试")
		}

		originUrl := in.GetOriginUrl()
		originUrl += uuid.New().String()
		shortUri = hashUtil.HashToBase62(originUrl)

		exists, _ := l.svcCtx.BloomFilter.Exists([]byte(shortUri))
		if !exists {
			// 布隆过滤器不存在的一定不存在 短链接唯一
			break
		}

		customGenerateCount++
	}

	return shortUri, nil
}

func (l *CreateShortLinkLogic) generateSuffixByLock(in *shortlink.ShortLinkCreateRequest) (string, error) {
	var (
		shortUri            string
		customGenerateCount = 0
	)

	for {
		if customGenerateCount > 10 {
			return "", errors.New("短链接频繁生成，请稍后再试")
		}

		originUrl := in.GetOriginUrl()
		originUrl += uuid.New().String()
		shortUri = hashUtil.HashToBase62(originUrl)

		_, err := l.svcCtx.ShortlinkModel.FindOneByShortUrl(l.ctx, shortUri, in.GetGid())
		if errors.Is(err, model.ErrNotFound) {
			break
		}
		customGenerateCount++
	}

	return shortUri, nil
}
