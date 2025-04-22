package shortlink

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	model "github.com/lambertstu/shortlink-core-rpc/mongo/shortlink"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

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
	// 保留现有统计数据
	TodayPv  string `json:"todayPv"`  // 今日PV
	TodayUv  string `json:"todayUv"`  // 今日UV
	TodayUip string `json:"todayUip"` // 今日IP
	TotalPv  string `json:"totalPv"`  // 总PV
	TotalUv  string `json:"totalUv"`  // 总UV
	TotalUip string `json:"totalUip"` // 总IP
	ClickNum string `json:"clickNum"` // 点击数
}

// 缓存过期时间设置
const (
	// 短链接缓存有效期 - 24小时
	shortLinkCacheExpiry = 24 * 60 * 60
	// 不存在的短链接缓存有效期 - 5分钟（避免缓存穿透）
	notExistCacheExpiry = 5 * 60
)

type RestoreUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestoreUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreUrlLogic {
	return &RestoreUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestoreUrlLogic) RestoreUrl(req *types.ShortLinkRestoreRequest) (resp *types.ShortLinkRestoreResponse, err error) {
	cacheKey := fmt.Sprintf("%s%s", l.svcCtx.ShortlinkRestoreKey, req.ShortUri)

	// 首先检查Redis缓存
	var shortlinkInfo *model.Shortlink
	var cacheHit bool

	// 第一次尝试从缓存获取数据
	jsonData, err := l.svcCtx.Redis.Get(cacheKey)
	if err == nil && jsonData != "" {
		// 缓存命中，解析JSON数据
		l.Logger.Infof("缓存命中: %s", cacheKey)
		cacheHit = true

		if jsonData == "NOT_EXIST" {
			l.Logger.Info("缓存显示此短链接不存在")
			return nil, errors.New("短链接不存在")
		}

		var cachedLink model.Shortlink
		err = json.Unmarshal([]byte(jsonData), &cachedLink)
		if err != nil {
			l.Logger.Errorf("解析缓存数据失败: %v", err)
		} else {
			shortlinkInfo = &cachedLink
		}
	}

	// 如果缓存未命中，先检查布隆过滤器
	if !cacheHit || shortlinkInfo == nil {
		// 检查布隆过滤器是否存在此短链接
		exists, err := l.svcCtx.BloomFilter.Exists([]byte(req.ShortUri))
		if err != nil {
			l.Logger.Errorf("检查布隆过滤器失败: %v", err)
		} else if !exists {
			// 布隆过滤器表明此短链接不存在，直接返回
			l.Logger.Infof("布隆过滤器显示短链接不存在: %s", req.ShortUri)

			return nil, errors.New("短链接不存在")
		}

		// 布隆过滤器表明可能存在，再次检查缓存
		// 这是为了处理并发情况下，其他请求可能已经重建了缓存
		l.Logger.Info("布隆过滤器显示短链接可能存在，二次检查缓存")
		jsonData, err = l.svcCtx.Redis.Get(cacheKey)
		if err == nil && jsonData != "" {
			// 二次缓存命中
			l.Logger.Info("二次缓存命中")

			if jsonData == "NOT_EXIST" {
				return nil, errors.New("短链接不存在")
			}

			var cachedLink model.Shortlink
			err = json.Unmarshal([]byte(jsonData), &cachedLink)
			if err == nil {
				shortlinkInfo = &cachedLink
			}
		}

		// 如果二次缓存仍未命中，尝试获取分布式锁并查询数据库
		if shortlinkInfo == nil {
			// 尝试获取分布式锁，防止缓存击穿
			acquiredLock, err := l.svcCtx.RedisLock.Acquire()
			if err != nil {
				l.Logger.Errorf("获取分布式锁失败: %v", err)
			}

			if acquiredLock {
				l.Logger.Info("获取分布式锁成功，查询数据库")
				defer func() {
					// 释放锁
					if _, err := l.svcCtx.RedisLock.Release(); err != nil {
						l.Logger.Errorf("释放分布式锁失败: %v", err)
					}
				}()

				// 获取锁成功后，再次检查缓存（可能其他线程已经构建好缓存）
				jsonData, err = l.svcCtx.Redis.Get(cacheKey)
				if err == nil && jsonData != "" && jsonData != "NOT_EXIST" {
					var cachedLink model.Shortlink
					if err := json.Unmarshal([]byte(jsonData), &cachedLink); err == nil {
						l.Logger.Info("获取锁后再次检查缓存成功")
						shortlinkInfo = &cachedLink
					}
				}

				// 如果仍未命中，查询数据库
				if shortlinkInfo == nil {
					l.Logger.Infof("缓存未命中，查询数据库: %s", req.ShortUri)
					var dbErr error
					shortlinkInfo, dbErr = l.svcCtx.ShortLinkModel.FindOneByShortUrl(l.ctx, req.ShortUri)

					if dbErr != nil {
						l.Logger.Errorf("查询数据库失败: %v", dbErr)

						if dbErr.Error() == "mongo: no documents in result" {
							// 设置不存在标记到缓存
							_ = l.svcCtx.Redis.SetexCtx(l.ctx, cacheKey, "NOT_EXIST", notExistCacheExpiry)
							l.Logger.Infof("已缓存不存在的短链接: %s, 有效期: %d秒", req.ShortUri, notExistCacheExpiry)
							return nil, errors.New("短链接不存在")
						}

						return nil, errors.New("短链接查询失败")
					}

					// 将数据库结果缓存到Redis
					linkJson, _ := json.Marshal(shortlinkInfo)
					err = l.svcCtx.Redis.SetexCtx(l.ctx, cacheKey, string(linkJson), shortLinkCacheExpiry)
					if err != nil {
						l.Logger.Errorf("缓存短链接数据失败: %v", err)
					} else {
						l.Logger.Infof("已缓存短链接数据: %s, 有效期: %d秒", req.ShortUri, shortLinkCacheExpiry)
					}

					// 将短链接添加到布隆过滤器
					if err := l.svcCtx.BloomFilter.Add([]byte(req.ShortUri)); err != nil {
						l.Logger.Errorf("添加到布隆过滤器失败: %v", err)
					}
				}
			} else {
				// 未获取到锁，等待一小段时间后重试缓存
				time.Sleep(20 * time.Millisecond)

				jsonData, err = l.svcCtx.Redis.Get(cacheKey)
				if err == nil && jsonData != "" {
					if jsonData == "NOT_EXIST" {
						return nil, errors.New("短链接不存在")
					}

					var cachedLink model.Shortlink
					if err := json.Unmarshal([]byte(jsonData), &cachedLink); err == nil {
						l.Logger.Info("等待后重试缓存成功")
						shortlinkInfo = &cachedLink
					} else {
						l.Logger.Errorf("未获取到锁且重试缓存失败: %v", err)
						return nil, errors.New("短链接查询失败")
					}
				} else {
					l.Logger.Error("未获取到锁且重试缓存未命中")
					return nil, errors.New("短链接查询失败，系统繁忙")
				}
			}
		}
	}

	// 构建点击数据
	clickData := ShortLinkClickData{
		Type:      "shortlink_click",
		Gid:       shortlinkInfo.Gid,
		ShortUri:  req.ShortUri,
		OriginUrl: shortlinkInfo.OriginUrl,
		Describe:  shortlinkInfo.Describe,
		Favicon:   shortlinkInfo.Favicon,
		ClickTime: time.Now(),
		TodayPv:   strconv.Itoa(shortlinkInfo.TodayPv),
		TodayUv:   strconv.Itoa(shortlinkInfo.TodayUv),
		TodayUip:  strconv.Itoa(shortlinkInfo.TodayUip),
		TotalPv:   strconv.Itoa(shortlinkInfo.TotalPv),
		TotalUv:   strconv.Itoa(shortlinkInfo.TotalUv),
		TotalUip:  strconv.Itoa(shortlinkInfo.TotalUip),
		ClickNum:  strconv.Itoa(shortlinkInfo.ClickNum),
	}

	// 将数据转为JSON并发送到消息队列（异步处理）
	go func() {
		// 创建新的上下文以避免使用主请求的上下文
		asyncCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		clickDataJson, err := json.Marshal(clickData)
		if err != nil {
			l.Logger.Errorf("序列化点击数据失败: %v", err)
			return
		}

		// 发送到消息队列
		queueName := "shortlink.v1.information.queue"
		err = l.svcCtx.RabbitmqSender.Send("", queueName, clickDataJson)
		if err != nil {
			// 使用异步上下文记录日志
			logx.WithContext(asyncCtx).Errorf("发送点击数据到消息队列失败: %v", err)
		} else {
			logx.WithContext(asyncCtx).Infof("已发送点击数据到消息队列: %s", queueName)
		}
	}()

	return &types.ShortLinkRestoreResponse{
		OriginUrl: shortlinkInfo.OriginUrl,
	}, nil
}
