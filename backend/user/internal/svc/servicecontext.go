package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"user/internal/config"
	groupModel "user/mongo/group"
	userModel "user/mongo/user"
)

const (
	mongoUrl = "mongodb://localhost:27019"
	bloomKey = "redis-bloom:user"
	bloomBit = 64
)

type ServiceContext struct {
	Config      config.Config
	Redis       *redis.Redis
	BloomFilter *bloom.Filter
	UserModel   userModel.UserModel
	GroupModel  groupModel.GroupModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, filter := getRedisInstance(c)
	return &ServiceContext{
		Config:      c,
		UserModel:   userModel.NewUserModel(mongoUrl, userModel.DB, userModel.Collection),
		Redis:       rds,
		BloomFilter: filter,
		GroupModel:  groupModel.NewGroupModel(mongoUrl, groupModel.DB, groupModel.Collection),
	}
}

func getRedisInstance(c config.Config) (*redis.Redis, *bloom.Filter) {
	rds := redis.MustNewRedis(c.RedisConf)
	filter := bloom.New(rds, bloomKey, bloomBit)
	return rds, filter
}
