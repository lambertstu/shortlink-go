package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	DB         = "shortlink"
	Collection = "shortlink"
)

var _ ShortlinkModel = (*customShortlinkModel)(nil)

type (
	// ShortlinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortlinkModel.
	ShortlinkModel interface {
		shortlinkModel
		FindOneByShortUrl(ctx context.Context, shortLink, gid string) (*Shortlink, error)
		InsertOneShortlink(ctx context.Context, data *Shortlink) error
	}

	customShortlinkModel struct {
		*defaultShortlinkModel
	}
)

func (c *customShortlinkModel) InsertOneShortlink(ctx context.Context, data *Shortlink) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	filter := bson.M{
		"short_uri":  data.ShortUri,
		"deleteFlag": bson.M{"$ne": 1},
	}

	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrShortUriExist
	}

	_, err = c.conn.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *customShortlinkModel) FindOneByShortUrl(ctx context.Context, shortLink, gid string) (*Shortlink, error) {
	filter := bson.M{
		"gid":        gid,
		"short_uri":  shortLink,
		"deleteFlag": 0,
	}

	var shortlink Shortlink

	err := c.conn.FindOne(ctx, &shortlink, filter)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &shortlink, nil
}

// NewShortlinkModel returns a model for the mongo.
func NewShortlinkModel(url, db, collection string) ShortlinkModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customShortlinkModel{
		defaultShortlinkModel: newDefaultShortlinkModel(conn),
	}
}
