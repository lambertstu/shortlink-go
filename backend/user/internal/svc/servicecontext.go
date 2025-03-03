package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"user/internal/config"
	user "user/mongo/user"
)

const mongoUrl = "mongodb://localhost:27019"

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(mongoUrl, user.DB, user.Collection),
		Redis:     redis.MustNewRedis(c.RedisConf),
	}
}
