package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"shortlink/internal/config"
	shortlink "shortlink/mongo/shortlink"
)

const (
	mongoUrl = "mongodb://localhost:27019"
	bloomKey = "redis-bloom:shortlink"
	bloomBit = 64
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redis.Redis
	BloomFilter    *bloom.Filter
	ShortlinkModel shortlink.ShortlinkModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, filter := getRedisInstance(c)
	return &ServiceContext{
		Config:         c,
		Redis:          rds,
		BloomFilter:    filter,
		ShortlinkModel: shortlink.NewShortlinkModel(mongoUrl, shortlink.DB, shortlink.Collection),
	}
}

func getRedisInstance(c config.Config) (*redis.Redis, *bloom.Filter) {
	rds := redis.MustNewRedis(c.RedisConf)
	filter := bloom.New(rds, bloomKey, bloomBit)
	return rds, filter
}
