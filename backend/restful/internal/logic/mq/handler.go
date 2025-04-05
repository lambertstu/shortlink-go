package mq

import (
	"encoding/json"
	"fmt"
	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

// MessageHandler 消息处理器接口
// 每种类型的消息处理逻辑需要实现此接口
type MessageHandler interface {
	// HandleMessage 处理具体消息
	HandleMessage(msgData map[string]interface{}) error
	// SupportedType 返回此处理器支持的消息类型
	SupportedType() string
}

// MQManager 消息队列管理器
// 用于管理所有消息处理器
type MQManager struct {
	CoreRpcService shortlinkclient.Shortlink
	handlers       map[string]MessageHandler
	mu             sync.RWMutex
	logger         logx.Logger
}

// NewMQManager 创建一个新的消息队列管理器
func NewMQManager(coreRpcService shortlinkclient.Shortlink) *MQManager {
	return &MQManager{
		CoreRpcService: coreRpcService,
		handlers:       make(map[string]MessageHandler),
		logger:         logx.WithContext(nil),
	}
}

// RegisterHandler 注册一个消息处理器
func (m *MQManager) RegisterHandler(handler MessageHandler) {
	m.mu.Lock()
	defer m.mu.Unlock()
	msgType := handler.SupportedType()
	m.handlers[msgType] = handler
	m.logger.Infof("注册消息处理器: %s", msgType)
}

// Handler 实现 ConsumeHandler 接口
// 用于分发消息到对应的处理器
type Handler struct {
	manager *MQManager
	logger  logx.Logger
}

// NewHandler 创建一个新的处理器分发器
func NewHandler(manager *MQManager) *Handler {
	return &Handler{
		manager: manager,
		logger:  logx.WithContext(nil),
	}
}

// Consume 消息处理方法
// 当从队列收到消息时会调用此方法
func (h *Handler) Consume(message string) error {
	h.logger.Infof("收到消息: %s", message)

	// 解析消息内容
	var msgData map[string]interface{}
	if err := json.Unmarshal([]byte(message), &msgData); err != nil {
		h.logger.Errorf("解析消息失败: %v", err)
		return err
	}

	// 获取消息类型
	msgType, ok := msgData["type"].(string)
	if !ok {
		h.logger.Error("消息缺少type字段或类型错误")
		return fmt.Errorf("消息缺少type字段或类型错误")
	}

	// 查找对应的处理器
	h.manager.mu.RLock()
	handler, exists := h.manager.handlers[msgType]
	h.manager.mu.RUnlock()

	if !exists {
		h.logger.Errorf("未找到处理消息类型 %s 的处理器", msgType)
		return fmt.Errorf("未找到处理消息类型 %s 的处理器", msgType)
	}

	// 调用对应处理器处理消息
	return handler.HandleMessage(msgData)
}

// 测试消息处理器示例
type TestMessageHandler struct {
	logger logx.Logger
}

// NewTestMessageHandler 创建测试消息处理器
func NewTestMessageHandler() *TestMessageHandler {
	return &TestMessageHandler{
		logger: logx.WithContext(nil),
	}
}

// HandleMessage 处理测试消息
func (h *TestMessageHandler) HandleMessage(msgData map[string]interface{}) error {
	message, _ := msgData["message"].(string)
	queue, _ := msgData["queue"].(string)

	h.logger.Infof("处理测试消息: %s, 来自队列: %s", message, queue)
	return nil
}

// SupportedType 返回支持的消息类型
func (h *TestMessageHandler) SupportedType() string {
	return "test"
}
