package mq

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/queue"
	"github.com/zeromicro/go-zero/core/threading"
)

// MQStarter RabbitMQ 启动器
// 用于非阻塞方式启动 RabbitMQ 监听器
type MQStarter struct {
	listeners []queue.MessageQueue
	logger    logx.Logger
}

// NewMQStarter 创建一个新的 MQ 启动器
func NewMQStarter() *MQStarter {
	return &MQStarter{
		listeners: make([]queue.MessageQueue, 0),
		logger:    logx.WithContext(nil),
	}
}

// AddListener 添加一个消息队列监听器
func (s *MQStarter) AddListener(listener queue.MessageQueue) {
	s.listeners = append(s.listeners, listener)
}

// Start 以非阻塞方式启动所有监听器
func (s *MQStarter) Start() {
	for i, listener := range s.listeners {
		// 使用独立的 goroutine 启动每个监听器
		// 避免阻塞主线程
		index := i
		l := listener
		threading.GoSafe(func() {
			s.logger.Infof("启动 RabbitMQ 监听器 [%d]", index)
			l.Start()
		})
	}

	// 注册优雅关闭钩子
	s.registerOnShutdown()
}

// Stop 停止所有监听器
func (s *MQStarter) Stop() {
	for i, listener := range s.listeners {
		s.logger.Infof("停止 RabbitMQ 监听器 [%d]", i)
		listener.Stop()
	}
}

// registerOnShutdown 注册程序关闭时的清理逻辑
func (s *MQStarter) registerOnShutdown() {
	proc.AddShutdownListener(func() {
		s.logger.Info("程序关闭，停止所有 RabbitMQ 监听器")
		s.Stop()
	})
}
