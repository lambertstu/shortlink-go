// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type shortlinkModel interface {
	Insert(ctx context.Context, data *Shortlink) error
	FindOne(ctx context.Context, id string) (*Shortlink, error)
	Update(ctx context.Context, data *Shortlink) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultShortlinkModel struct {
	conn *mon.Model
}

func newDefaultShortlinkModel(conn *mon.Model) *defaultShortlinkModel {
	return &defaultShortlinkModel{conn: conn}
}

func (m *defaultShortlinkModel) Insert(ctx context.Context, data *Shortlink) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultShortlinkModel) FindOne(ctx context.Context, id string) (*Shortlink, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Shortlink

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortlinkModel) Update(ctx context.Context, data *Shortlink) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultShortlinkModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
