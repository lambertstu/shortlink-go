package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	DB         = "shortlink"
	Collection = "user"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		Register(ctx context.Context, data *User) error
		UpdateUserInfo(ctx context.Context, data *User) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (c *customUserModel) Register(ctx context.Context, data *User) error {
	filter := bson.M{"username": data.Username}

	// TODO 可以引入布隆过滤器优化
	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}

	_, err = c.conn.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *customUserModel) UpdateUserInfo(ctx context.Context, data *User) error {
	filter := bson.D{
		{"username", data.Username},
	}
	update := bson.D{
		{"$set", data},
	}
	_, err := c.conn.UpdateOne(ctx, filter, update)
	return err
}

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string) UserModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}
