package rabbitmq

import "fmt"

// RabbitConf RabbitMQ基础配置结构体
// 定义连接RabbitMQ服务器所需的基本信息
type RabbitConf struct {
	Username string // 用户名
	Password string // 密码
	Host     string // 主机地址
	Port     int    // 端口号
	VHost    string `json:",optional"` // 虚拟主机，可选参数
}

// RabbitListenerConf 消息消费者配置结构体
// 继承基础配置并增加消费者队列配置
type RabbitListenerConf struct {
	RabbitConf                    // 嵌入基础配置
	ListenerQueues []ConsumerConf // 消费者队列配置列表
}

// ConsumerConf 消费者详细配置结构体
type ConsumerConf struct {
	Name      string // 队列名称
	AutoAck   bool   `json:",default=true"`  // 自动确认消息，默认为true
	Exclusive bool   `json:",default=false"` // 排他性，默认为false，设为true时队列只对首次声明它的连接可见
	// NoLocal 设置为true时，表示该连接的生产者发送的消息不能被同一连接的消费者消费
	NoLocal bool `json:",default=false"`
	// NoWait 是否阻塞处理
	NoWait bool `json:",default=false"` // 是否等待服务器响应，默认为false
}

// RabbitSenderConf 消息生产者配置结构体
type RabbitSenderConf struct {
	RabbitConf         // 嵌入基础配置
	ContentType string `json:",default=text/plain"` // MIME内容类型，默认为纯文本
}

// QueueConf 队列配置结构体
type QueueConf struct {
	Name       string // 队列名称
	Durable    bool   `json:",default=true"`  // 是否持久化，默认为true
	AutoDelete bool   `json:",default=false"` // 是否自动删除，默认为false，当最后一个消费者断开连接后队列是否自动删除
	Exclusive  bool   `json:",default=false"` // 是否排他性，默认为false
	NoWait     bool   `json:",default=false"` // 是否等待服务器响应，默认为false
}

// ExchangeConf 交换机配置结构体
type ExchangeConf struct {
	ExchangeName string      // 交换机名称
	Type         string      `json:",options=direct|fanout|topic|headers"` // 交换机类型：直接交换机|扇出交换机|主题交换机|头交换机
	Durable      bool        `json:",default=true"`                        // 是否持久化，默认为true
	AutoDelete   bool        `json:",default=false"`                       // 是否自动删除，默认为false
	Internal     bool        `json:",default=false"`                       // 是否为内部交换机，默认为false，设为true时交换机只能被RabbitMQ内部使用
	NoWait       bool        `json:",default=false"`                       // 是否等待服务器响应，默认为false
	Queues       []QueueConf // 与该交换机绑定的队列配置列表
}

// getRabbitURL 根据RabbitMQ配置生成连接URL
// 返回格式为amqp://username:password@host:port/vhost的连接字符串
func getRabbitURL(rabbitConf RabbitConf) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rabbitConf.Username, rabbitConf.Password,
		rabbitConf.Host, rabbitConf.Port, rabbitConf.VHost)
}
