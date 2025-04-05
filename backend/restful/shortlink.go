package main

import (
	"flag"
	"fmt"

	"github.com/lambertstu/shortlink-go/restful/internal/config"
	"github.com/lambertstu/shortlink-go/restful/internal/handler"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shortlink-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// rabbitmq配置初始化
	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(ctx.RabbitmqListener)
	defer serviceGroup.Stop()

	// 使用 goroutine 启动 RabbitMQ 监听器
	go serviceGroup.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
