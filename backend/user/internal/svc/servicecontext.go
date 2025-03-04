package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"user/internal/config"
	user "user/mongo/user"
)

const (
	mongoUrl = "mongodb://localhost:27019"
	bloomKey = "redis-bloom:user"
	bloomBit = 64
)

type ServiceContext struct {
	Config      config.Config
	UserModel   user.UserModel
	Redis       *redis.Redis
	BloomFilter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	redis, filter := getRedisInstance(c)
	return &ServiceContext{
		Config:      c,
		UserModel:   user.NewUserModel(mongoUrl, user.DB, user.Collection),
		Redis:       redis,
		BloomFilter: filter,
	}
}

func getRedisInstance(c config.Config) (*redis.Redis, *bloom.Filter) {
	rds := redis.MustNewRedis(c.RedisConf)
	filter := bloom.New(rds, bloomKey, bloomBit)
	return rds, filter
}
