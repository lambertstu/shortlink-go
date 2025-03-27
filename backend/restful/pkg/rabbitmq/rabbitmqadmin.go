package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Admin RabbitMQ管理结构体
// 提供交换机、队列的声明和绑定等管理功能
type Admin struct {
	conn    *amqp.Connection // RabbitMQ连接
	channel *amqp.Channel    // RabbitMQ通道
}

// MustNewAdmin 创建并初始化一个新的RabbitMQ管理器
// 如果初始化过程中出现错误，会直接导致程序崩溃
// 参数:
//   - rabbitMqConf: RabbitMQ基础配置
//
// 返回: RabbitMQ管理器指针
func MustNewAdmin(rabbitMqConf RabbitConf) *Admin {
	var admin Admin
	conn, err := amqp.Dial(getRabbitURL(rabbitMqConf))
	if err != nil {
		// 连接RabbitMQ服务器失败，记录错误并终止程序
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}

	admin.conn = conn
	channel, err := admin.conn.Channel()
	if err != nil {
		// 打开通道失败，记录错误并终止程序
		log.Fatalf("failed to open a channel, error: %v", err)
	}

	admin.channel = channel
	return &admin
}

// DeclareExchange 声明一个交换机
// 参数:
//   - conf: 交换机配置
//   - args: 额外参数表
//
// 返回: 错误信息，如果成功则为nil
func (q *Admin) DeclareExchange(conf ExchangeConf, args amqp.Table) error {
	return q.channel.ExchangeDeclare(
		conf.ExchangeName, // 交换机名称
		conf.Type,         // 交换机类型（direct, fanout, topic, headers）
		conf.Durable,      // 是否持久化
		conf.AutoDelete,   // 是否自动删除
		conf.Internal,     // 是否为内部交换机
		conf.NoWait,       // 是否等待服务器响应
		args,              // 额外参数
	)
}

// DeclareQueue 声明一个队列
// 参数:
//   - conf: 队列配置
//   - args: 额外参数表
//
// 返回: 错误信息，如果成功则为nil
func (q *Admin) DeclareQueue(conf QueueConf, args amqp.Table) error {
	_, err := q.channel.QueueDeclare(
		conf.Name,       // 队列名称
		conf.Durable,    // 是否持久化
		conf.AutoDelete, // 是否自动删除
		conf.Exclusive,  // 是否排他
		conf.NoWait,     // 是否等待服务器响应
		args,            // 额外参数
	)

	return err
}

// Bind 将队列绑定到交换机
// 参数:
//   - queueName: 队列名称
//   - routekey: 路由键
//   - exchange: 交换机名称
//   - notWait: 是否等待服务器响应
//   - args: 额外参数表
//
// 返回: 错误信息，如果成功则为nil
func (q *Admin) Bind(queueName string, routekey string, exchange string, notWait bool, args amqp.Table) error {
	return q.channel.QueueBind(
		queueName, // 队列名称
		routekey,  // 路由键，用于消息路由到队列
		exchange,  // 交换机名称
		notWait,   // 是否等待服务器响应
		args,      // 额外参数
	)
}
