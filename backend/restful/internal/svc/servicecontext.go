package svc

import (
	model "github.com/lambertstu/shortlink-core-rpc/mongo/shortlink"
	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"
	"github.com/lambertstu/shortlink-go/restful/internal/config"
	"github.com/lambertstu/shortlink-go/restful/internal/logic/mq"
	"github.com/lambertstu/shortlink-go/restful/pkg/rabbitmq"
	"github.com/lambertstu/shortlink-user-rpc/client/group"
	"github.com/lambertstu/shortlink-user-rpc/client/user"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	bloomKey            = "redis-bloom:shortlink-restore"
	bloomBit            = 64
	ShortlinkRestoreKey = "shortlink-restore:"
)

type ServiceContext struct {
	Config              config.Config
	UserRpcService      user.User
	GroupRpcService     group.Group
	CoreRpcService      shortlinkclient.Shortlink
	ShortLinkModel      model.ShortlinkModel
	RabbitmqListener    queue.MessageQueue
	RabbitmqSender      rabbitmq.Sender
	MQManager           *mq.MQManager
	Redis               *redis.Redis
	BloomFilter         *bloom.Filter
	ShortlinkRestoreKey string
}

func NewServiceContext(c config.Config) *ServiceContext {
	coreRpcService := shortlinkclient.NewShortlink(zrpc.MustNewClient(c.CoreRpcConf))

	mqManager := mq.NewMQManager(coreRpcService)
	mqManager.RegisterHandler(mq.NewTestMessageHandler())
	mqManager.RegisterHandler(mq.NewShortLinkClickHandler(coreRpcService))

	mqHandler := mq.NewHandler(mqManager)

	rds, filter := getRedisInstance(c)

	return &ServiceContext{
		Config:              c,
		UserRpcService:      user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		GroupRpcService:     group.NewGroup(zrpc.MustNewClient(c.GroupRpcConf)),
		CoreRpcService:      coreRpcService,
		ShortLinkModel:      model.NewShortlinkModel(c.ShortLinkModelUrl, model.DB, model.Collection),
		RabbitmqListener:    rabbitmq.MustNewListener(c.ListenerConf, mqHandler),
		RabbitmqSender:      rabbitmq.MustNewSender(c.SenderConf),
		MQManager:           mqManager,
		Redis:               rds,
		BloomFilter:         filter,
		ShortlinkRestoreKey: ShortlinkRestoreKey,
	}
}

func getRedisInstance(c config.Config) (*redis.Redis, *bloom.Filter) {
	rds := redis.MustNewRedis(c.RedisConf.RedisConf)
	filter := bloom.New(rds, bloomKey, bloomBit)
	return rds, filter
}
