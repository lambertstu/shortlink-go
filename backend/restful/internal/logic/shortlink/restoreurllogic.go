package shortlink

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

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
	shortlinkInfo, err := l.svcCtx.ShortLinkModel.FindOneByShortUrl(l.ctx, req.ShortUri)
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("短链接查询失败")
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

	// 将数据转为JSON
	clickDataJson, err := json.Marshal(clickData)
	if err != nil {
		l.Logger.Errorf("序列化点击数据失败: %v", err)
	} else {
		// 发送到消息队列
		queueName := "shortlink.v1.information.queue"
		err = l.svcCtx.RabbitmqSender.Send("", queueName, clickDataJson)
		if err != nil {
			l.Logger.Errorf("发送点击数据到消息队列失败: %v", err)
		} else {
			l.Logger.Infof("已发送点击数据到消息队列: %s", queueName)
		}
	}

	return &types.ShortLinkRestoreResponse{
		OriginUrl: shortlinkInfo.OriginUrl,
	}, nil
}
