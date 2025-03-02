package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string             `bson:"username,omitempty" json:"username,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	RealName     string             `bson:"realName,omitempty" json:"realName,omitempty"`
	Phone        string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Mail         string             `bson:"mail,omitempty" json:"mail,omitempty"`
	DeleteFlag   int                `bson:"deleteFlag" json:"deleteFlag,omitempty"`
	DeletionTime time.Time          `bson:"deletionTime,omitempty" json:"deletionTime,omitempty"`
	UpdateAt     time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
