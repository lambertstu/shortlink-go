package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GID      string             `bson:"gid,omitempty" json:"gid,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`         // 分组名
	Username string             `bson:"username,omitempty" json:"username,omitempty"` // 创建分组的用户名
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
