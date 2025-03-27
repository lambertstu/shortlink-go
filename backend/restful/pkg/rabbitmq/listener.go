package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
)

type (
	// ConsumeHandle 消息处理函数类型定义
	// 接收一个字符串消息并返回错误
	ConsumeHandle func(message string) error

	// ConsumeHandler 消息处理接口
	// 定义了消息消费的行为
	ConsumeHandler interface {
		Consume(message string) error
	}

	// RabbitListener RabbitMQ监听器结构体
	// 用于从RabbitMQ队列接收和处理消息
	RabbitListener struct {
		conn    *amqp.Connection   // RabbitMQ连接
		channel *amqp.Channel      // RabbitMQ通道
		forever chan bool          // 用于保持监听器运行的通道
		handler ConsumeHandler     // 消息处理器
		queues  RabbitListenerConf // 监听器配置
	}
)

// MustNewListener 创建并初始化一个新的RabbitMQ监听器
// 如果初始化过程中出现错误，会直接导致程序崩溃
// 参数:
//   - listenerConf: 监听器配置
//   - handler: 消息处理器
//
// 返回: 消息队列接口
func MustNewListener(listenerConf RabbitListenerConf, handler ConsumeHandler) queue.MessageQueue {
	listener := RabbitListener{queues: listenerConf, handler: handler, forever: make(chan bool)}
	conn, err := amqp.Dial(getRabbitURL(listenerConf.RabbitConf))
	if err != nil {
		// 连接RabbitMQ服务器失败，记录错误并终止程序
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}

	listener.conn = conn
	channel, err := listener.conn.Channel()
	if err != nil {
		// 打开通道失败，记录错误并终止程序
		log.Fatalf("failed to open a channel: %v", err)
	}

	listener.channel = channel
	return listener
}

// Start 开始监听RabbitMQ队列中的消息
// 为每个配置的队列启动一个goroutine来消费消息
func (q RabbitListener) Start() {
	for _, que := range q.queues.ListenerQueues {
		// 为每个队列设置消费者
		msg, err := q.channel.Consume(
			que.Name,      // 队列名称
			"",            // 消费者标签，空字符串表示自动生成
			que.AutoAck,   // 是否自动确认消息
			que.Exclusive, // 是否独占队列
			que.NoLocal,   // 是否不接收同一连接生产的消息
			que.NoWait,    // 是否等待服务器响应
			nil,           // 额外参数
		)
		if err != nil {
			// 设置消费者失败，记录错误并终止程序
			log.Fatalf("failed to listener, error: %v", err)
		}

		// 为每个队列启动一个goroutine来处理消息
		go func() {
			for d := range msg {
				// 调用处理器处理消息
				if err := q.handler.Consume(string(d.Body)); err != nil {
					// 处理消息失败，记录错误但继续处理下一条消息
					logx.Errorf("Error on consuming: %s, error: %v", string(d.Body), err)
				}
			}
		}()
	}

	// 阻塞主线程，保持监听器运行
	<-q.forever
}

// Stop 停止监听器
// 关闭通道和连接，释放资源
func (q RabbitListener) Stop() {
	q.channel.Close()
	q.conn.Close()
	close(q.forever)
}
