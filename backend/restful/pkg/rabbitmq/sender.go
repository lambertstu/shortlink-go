package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type (
	// Sender 消息发送器接口
	// 定义了发送消息到交换机的行为
	Sender interface {
		Send(exchange string, routeKey string, msg []byte) error
	}

	// RabbitMqSender RabbitMQ发送器结构体
	// 实现了Sender接口，用于向RabbitMQ交换机发送消息
	RabbitMqSender struct {
		conn        *amqp.Connection // RabbitMQ连接
		channel     *amqp.Channel    // RabbitMQ通道
		ContentType string           // 消息内容类型
	}
)

// MustNewSender 创建并初始化一个新的RabbitMQ发送器
// 如果初始化过程中出现错误，会直接导致程序崩溃
// 参数:
//   - rabbitMqConf: RabbitMQ发送器配置
//
// 返回: 发送器接口
func MustNewSender(rabbitMqConf RabbitSenderConf) Sender {
	sender := &RabbitMqSender{ContentType: rabbitMqConf.ContentType}
	conn, err := amqp.Dial(getRabbitURL(rabbitMqConf.RabbitConf))
	if err != nil {
		// 连接RabbitMQ服务器失败，记录错误并终止程序
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}

	sender.conn = conn
	channel, err := sender.conn.Channel()
	if err != nil {
		// 打开通道失败，记录错误并终止程序
		log.Fatalf("failed to open a channel, error: %v", err)
	}

	sender.channel = channel
	return sender
}

// Send 发送消息到指定的交换机和路由键
// 参数:
//   - exchange: 交换机名称
//   - routeKey: 路由键
//   - msg: 消息内容（字节数组）
//
// 返回: 错误信息，如果成功则为nil
func (q *RabbitMqSender) Send(exchange string, routeKey string, msg []byte) error {
	return q.channel.PublishWithContext(
		context.Background(), // 上下文对象，用于控制发布过程
		exchange,             // 交换机名称
		routeKey,             // 路由键，用于决定消息被路由到哪个队列
		false,                // mandatory，如果为true，消息不能路由到队列时会返回给发送者
		false,                // immediate，如果为true，消息不能立即投递到消费者时会返回给发送者
		amqp.Publishing{
			ContentType: q.ContentType, // 消息内容类型
			Body:        msg,           // 消息内容
		},
	)
}
