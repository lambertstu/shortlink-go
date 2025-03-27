package svc

import (
	"github.com/lambertstu/shortlink-core-rpc/mongo/shortlink"
	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"
	"github.com/lambertstu/shortlink-go/restful/internal/config"
	"github.com/lambertstu/shortlink-go/restful/internal/logic/mq"
	"github.com/lambertstu/shortlink-go/restful/pkg/rabbitmq"
	"github.com/lambertstu/shortlink-user-rpc/client/group"
	"github.com/lambertstu/shortlink-user-rpc/client/user"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	UserRpcService   user.User
	GroupRpcService  group.Group
	CoreRpcService   shortlinkclient.Shortlink
	ShortLinkModel   model.ShortlinkModel
	RabbitmqListener queue.MessageQueue
	RabbitmqSender   rabbitmq.Sender
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserRpcService:   user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		GroupRpcService:  group.NewGroup(zrpc.MustNewClient(c.GroupRpcConf)),
		CoreRpcService:   shortlinkclient.NewShortlink(zrpc.MustNewClient(c.CoreRpcConf)),
		ShortLinkModel:   model.NewShortlinkModel(c.ShortLinkModelUrl, model.DB, model.Collection),
		RabbitmqListener: rabbitmq.MustNewListener(c.ListenerConf, mq.Handler{}),
		RabbitmqSender:   rabbitmq.MustNewSender(c.SenderConf),
	}
}
