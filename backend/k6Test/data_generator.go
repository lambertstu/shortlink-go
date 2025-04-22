package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ShortLink 数据结构
type ShortLink struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Domain       string             `bson:"domain"`
	OriginUrl    string             `bson:"origin_url"`
	Gid          string             `bson:"gid"`
	Describe     string             `bson:"describe"`
	FullShortUrl string             `bson:"full_short_url"`
	ShortUri     string             `bson:"short_uri"`
	Favicon      string             `bson:"favicon"`
	ClickNum     int                `bson:"click_num"`
	TotalPv      int                `bson:"total_pv"`
	TotalUv      int                `bson:"total_uv"`
	TotalUip     int                `bson:"total_uip"`
	TodayPv      int                `bson:"today_pv"`
	TodayUv      int                `bson:"today_uv"`
	TodayUip     int                `bson:"today_uip"`
	DeleteFlag   int                `bson:"deleteFlag"`
	UpdateAt     primitive.DateTime `bson:"updateAt"`
	CreateAt     primitive.DateTime `bson:"createAt"`
}

// 随机字符串生成器
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// 生成随机短链接
func generateRandomShortLink() ShortLink {
	// 随机生成6位短链接URI
	shortUri := randomString(6)

	// 随机选择一个源网站
	websites := []string{
		"https://www.bilibili.com/video/BV" + randomString(10),
		"https://www.youtube.com/watch?v=" + randomString(11),
		"https://github.com/" + randomString(8) + "/" + randomString(10),
		"https://www.zhihu.com/question/" + fmt.Sprintf("%d", rand.Intn(999999999)),
		"https://mp.weixin.qq.com/s/" + randomString(16),
	}
	originUrl := websites[rand.Intn(len(websites))]

	// 随机生成描述
	descriptions := []string{
		"视频分享链接",
		"技术博客",
		"学习资料",
		"产品介绍",
		"新闻资讯",
		"官方文档",
		"精彩内容",
	}

	// 随机生成未来的时间
	futureDate := time.Now().AddDate(0, rand.Intn(24), rand.Intn(30))

	// 随机生成过去的时间
	pastDate := time.Now().AddDate(0, -rand.Intn(6), -rand.Intn(30))

	// 生成一个4位随机字符作为GID
	gid := randomString(4)

	// 生成完整短链接
	fullShortUrl := "zeroLink:8001/" + shortUri

	clickNum := rand.Intn(1000)
	totalPv := clickNum + rand.Intn(100)
	totalUv := totalPv / 2
	totalUip := totalUv / 2

	// 今日数据是总数据的一小部分
	todayPv := 0
	if totalPv > 10 {
		todayPv = rand.Intn(totalPv / 10)
	}

	todayUv := 0
	if totalUv > 10 {
		todayUv = rand.Intn(totalUv / 10)
	}

	todayUip := 0
	if totalUip > 10 {
		todayUip = rand.Intn(totalUip / 10)
	}

	return ShortLink{
		Domain:       "nurl.ink",
		OriginUrl:    originUrl,
		Gid:          gid,
		Describe:     descriptions[rand.Intn(len(descriptions))],
		FullShortUrl: fullShortUrl,
		ShortUri:     shortUri,
		Favicon:      "http/example.com",
		ClickNum:     clickNum,
		TotalPv:      totalPv,
		TotalUv:      totalUv,
		TotalUip:     totalUip,
		TodayPv:      todayPv,
		TodayUv:      todayUv,
		TodayUip:     todayUip,
		DeleteFlag:   0,
		UpdateAt:     primitive.NewDateTimeFromTime(futureDate),
		CreateAt:     primitive.NewDateTimeFromTime(pastDate),
	}
}

func main() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 设置MongoDB连接
	mongoURI := "mongodb://localhost:27019"
	clientOptions := options.Client().ApplyURI(mongoURI)

	// 连接MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("连接MongoDB失败: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("断开MongoDB连接失败: %v", err)
		}
	}()

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("无法ping MongoDB: %v", err)
	}

	fmt.Println("成功连接到MongoDB!")

	// 获取数据库和集合
	db := client.Database("shortlink")
	collection := db.Collection("shortlink")

	// 创建短链接URI的唯一索引
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "short_uri", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("创建索引失败: %v", err)
	}

	// 批量插入参数
	batchSize := 1000
	totalRecords := 50000
	var successCount int
	var failCount int
	var duplicateCount int

	fmt.Printf("开始插入%d条记录，批次大小: %d\n", totalRecords, batchSize)
	startTime := time.Now()

	// 分批次插入数据
	for i := 0; i < totalRecords; i += batchSize {
		var batch []interface{}
		currentBatchSize := batchSize
		if i+batchSize > totalRecords {
			currentBatchSize = totalRecords - i
		}

		// 构建批次数据
		for j := 0; j < currentBatchSize; j++ {
			shortlink := generateRandomShortLink()
			batch = append(batch, shortlink)
		}

		// 插入记录
		insertResult, err := collection.InsertMany(context.Background(), batch)
		if err != nil {
			// 检查是否是部分写入错误
			if strings.Contains(err.Error(), "duplicate key error") {
				log.Printf("批次 %d-%d: 存在重复键", i, i+currentBatchSize-1)
				// 当出现重复键错误时，insertResult可能为nil
				if insertResult != nil {
					duplicateCount += batchSize - len(insertResult.InsertedIDs)
					successCount += len(insertResult.InsertedIDs)
				} else {
					duplicateCount += currentBatchSize
				}
			} else {
				log.Printf("批次 %d-%d: 插入失败: %v", i, i+currentBatchSize-1, err)
				failCount += currentBatchSize
			}
		} else {
			log.Printf("批次 %d-%d: 成功插入 %d 条记录", i, i+currentBatchSize-1, len(insertResult.InsertedIDs))
			successCount += len(insertResult.InsertedIDs)
		}

		// 显示进度
		progress := float64(i+currentBatchSize) / float64(totalRecords) * 100
		fmt.Printf("已完成: %.2f%% (%d/%d)\n", progress, i+currentBatchSize, totalRecords)
	}

	// 显示统计信息
	duration := time.Since(startTime)
	fmt.Printf("\n操作完成!\n")
	fmt.Printf("总耗时: %v\n", duration)
	fmt.Printf("成功插入: %d 条记录\n", successCount)
	fmt.Printf("失败记录: %d 条\n", failCount)
	fmt.Printf("重复记录: %d 条\n", duplicateCount)
	fmt.Printf("平均插入速度: %.2f 条/秒\n", float64(successCount)/duration.Seconds())

	// 查询并显示数据库中的总记录数
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Printf("无法获取集合中的记录数: %v", err)
	} else {
		fmt.Printf("集合中现有记录数: %d\n", count)
	}

	fmt.Println("程序执行完毕")
	os.Exit(0)
}
