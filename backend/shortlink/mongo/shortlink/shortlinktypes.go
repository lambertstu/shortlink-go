package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shortlink struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Domain        string             `bson:"domain" json:"domain"`
	ShortUri      string             `bson:"short_uri" json:"shortUri"`
	FullShortUrl  string             `bson:"full_short_url" json:"fullShortUrl"`
	OriginUrl     string             `bson:"origin_url" json:"originUrl"`
	ClickNum      int                `bson:"click_num" json:"clickNum"`
	Gid           string             `bson:"gid" json:"gid"`
	EnableStatus  int                `bson:"enable_status" json:"enableStatus"`
	ValidDateType int32              `bson:"valid_date_type" json:"validDateType"`
	ValidDate     time.Time          `bson:"valid_date" json:"validDate"`
	Describe      string             `bson:"describe" json:"describe"`
	Favicon       string             `bson:"favicon" json:"favicon"`
	TotalPv       int                `bson:"total_pv" json:"totalPv"`
	TotalUv       int                `bson:"total_uv" json:"totalUv"`
	TotalUip      int                `bson:"total_uip" json:"totalUip"`
	TodayPv       int                `bson:"-" json:"todayPv,omitempty"`
	TodayUv       int                `bson:"-" json:"todayUv,omitempty"`
	TodayUip      int                `bson:"-" json:"todayUip,omitempty"`
	DeleteFlag    int                `bson:"deleteFlag" json:"deleteFlag"`
	DelTime       int64              `bson:"del_time" json:"delTime"`
	UpdateAt      time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt      time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
