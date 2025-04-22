package config

import (
	"github.com/lambertstu/shortlink-go/restful/pkg/rabbitmq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf       zrpc.RpcClientConf
	GroupRpcConf      zrpc.RpcClientConf
	CoreRpcConf       zrpc.RpcClientConf
	ShortLinkModelUrl string
	ListenerConf      rabbitmq.RabbitListenerConf
	SenderConf        rabbitmq.RabbitSenderConf
	RedisConf         RedisConf
}

type RedisConf struct {
	redis.RedisConf
}
