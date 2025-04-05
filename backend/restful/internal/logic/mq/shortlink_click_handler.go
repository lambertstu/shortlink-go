package mq

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/lambertstu/shortlink-core-rpc/shortlinkclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// ShortLinkClickData 短链接点击数据结构
type ShortLinkClickData struct {
	Type      string    `json:"type"`       // 消息类型
	Gid       string    `json:"gid"`        // 短链接标识
	ShortUri  string    `json:"short_uri"`  // 短链接URI
	OriginUrl string    `json:"origin_url"` // 原始URL
	Describe  string    `json:"describe"`   // 描述
	Favicon   string    `json:"favicon"`    // 图标
	ClickTime time.Time `json:"click_time"` // 点击时间
	// 统计数据
	TodayPv  string `json:"todayPv"`  // 今日PV
	TodayUv  string `json:"todayUv"`  // 今日UV
	TodayUip string `json:"todayUip"` // 今日IP
	TotalPv  string `json:"totalPv"`  // 总PV
	TotalUv  string `json:"totalUv"`  // 总UV
	TotalUip string `json:"totalUip"` // 总IP
	ClickNum string `json:"clickNum"` // 点击数
}

// ShortLinkClickHandler 短链接点击处理器
type ShortLinkClickHandler struct {
	logger         logx.Logger
	coreRpcService shortlinkclient.Shortlink
	// 跟踪当前日期
	currentDate string
	// 记录每个短链接最后处理的日期
	shortLinkLastProcessed map[string]string
	mu                     sync.RWMutex
}

// NewShortLinkClickHandler 创建新的短链接点击处理器
func NewShortLinkClickHandler(coreRpcService shortlinkclient.Shortlink) *ShortLinkClickHandler {
	// 获取当前日期作为初始日期
	currentDate := time.Now().Format("2006-01-02")

	return &ShortLinkClickHandler{
		logger:                 logx.WithContext(nil),
		coreRpcService:         coreRpcService,
		currentDate:            currentDate,
		shortLinkLastProcessed: make(map[string]string),
	}
}

// HandleMessage 处理短链接点击消息
func (h *ShortLinkClickHandler) HandleMessage(msgData map[string]interface{}) error {
	h.logger.Infof("处理短链接点击消息: %v", msgData)

	// 解析消息数据
	var clickData ShortLinkClickData
	jsonData, err := json.Marshal(msgData)
	if err != nil {
		h.logger.Errorf("序列化消息数据失败: %v", err)
		return err
	}

	if err := json.Unmarshal(jsonData, &clickData); err != nil {
		h.logger.Errorf("解析点击数据失败: %v", err)
		return err
	}

	// 获取当前日期
	currentDate := time.Now().Format("2006-01-02")

	// 检查是否是新的一天
	if currentDate != h.currentDate {
		h.mu.Lock()
		h.currentDate = currentDate
		h.mu.Unlock()
	}

	// 检查这个短链接今天是否已经处理过
	h.mu.RLock()
	lastProcessedDate, exists := h.shortLinkLastProcessed[clickData.ShortUri]
	h.mu.RUnlock()

	// 判断是否是今日第一条消息
	isFirstMessageOfDay := !exists || lastProcessedDate != currentDate

	// 获取现有数据
	todayPv, _ := strconv.Atoi(clickData.TodayPv)
	todayUv, _ := strconv.Atoi(clickData.TodayUv)
	todayUip, _ := strconv.Atoi(clickData.TodayUip)
	totalPv, _ := strconv.Atoi(clickData.TotalPv)
	totalUv, _ := strconv.Atoi(clickData.TotalUv)
	totalUip, _ := strconv.Atoi(clickData.TotalUip)
	clickNum, _ := strconv.Atoi(clickData.ClickNum)

	// 如果是今日第一条消息
	if isFirstMessageOfDay {
		h.logger.Infof("收到短链接 %s 今日第一条点击消息", clickData.ShortUri)
		// 将当日计数重置为1
		todayPv = 1
		todayUv = 1
		todayUip = 1

		// 记录该短链接已在今日处理过
		h.mu.Lock()
		h.shortLinkLastProcessed[clickData.ShortUri] = currentDate
		h.mu.Unlock()
	} else {
		// 非今日第一条消息，递增计数
		todayPv++
		todayUv++
		todayUip++
	}

	// 总量始终递增
	totalPv++
	totalUv++
	totalUip++
	clickNum++

	ctx := context.Background()

	_, err = h.coreRpcService.UpdateShortLink(ctx, &shortlinkclient.ShortLinkUpdateRequest{
		Gid:       clickData.Gid,
		OriginUrl: clickData.OriginUrl,
		ShortUri:  clickData.ShortUri,
		Describe:  clickData.Describe,
		Favicon:   clickData.Favicon,
		TodayPv:   int64(todayPv),
		TodayUip:  int64(todayUip),
		TodayUv:   int64(todayUv),
		TotalPv:   int64(totalPv),
		TotalUip:  int64(totalUip),
		TotalUv:   int64(totalUv),
		ClickNum:  int64(clickNum),
	})

	if err != nil {
		h.logger.Errorf("更新短链接 %s 统计数据失败: %v", clickData.ShortUri, err)
		return err
	}

	h.logger.Infof("成功更新短链接 %s 统计数据", clickData.ShortUri)
	return nil
}

// SupportedType 返回支持的消息类型
func (h *ShortLinkClickHandler) SupportedType() string {
	return "shortlink_click"
}
