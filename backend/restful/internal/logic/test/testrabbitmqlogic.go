package test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestRabbitMQLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestRabbitMQLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestRabbitMQLogic {
	return &TestRabbitMQLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestRabbitMQLogic) TestRabbitMQ(req *types.RabbitMQTestRequest) (resp *types.RabbitMQTestResponse, err error) {
	resp = &types.RabbitMQTestResponse{}

	queueName := "shortlink.v1.information.queue"
	if len(req.Queue) > 0 {
		queueName = req.Queue
	}

	// 将消息内容转换为 JSON 格式
	msgData := map[string]interface{}{
		"message": req.Message,
		"type":    "test",
		"queue":   queueName,
	}

	jsonData, err := json.Marshal(msgData)
	if err != nil {
		l.Error(fmt.Sprintf("序列化消息失败: %v", err))
		resp.Status = "error"
		resp.Message = fmt.Sprintf("序列化消息失败: %v", err)
		return resp, err
	}

	// 使用 RabbitMQ 发送消息
	// 参数: exchange(交换机名称), routeKey(路由键), msg(消息内容)
	// 这里我们使用默认交换机 ""，路由键为队列名称
	err = l.svcCtx.RabbitmqSender.Send("", queueName, jsonData)
	if err != nil {
		l.Error(fmt.Sprintf("发送消息到 RabbitMQ 失败: %v", err))
		resp.Status = "error"
		resp.Message = fmt.Sprintf("发送消息到 RabbitMQ 失败: %v", err)
		return resp, err
	}

	// 消息发送成功
	resp.Status = "success"
	resp.Message = fmt.Sprintf("消息已成功发送到队列 %s: %s", queueName, req.Message)

	return resp, nil
}
