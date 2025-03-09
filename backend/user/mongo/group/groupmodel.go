package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"user/pb/user"
)

const (
	DB         = "shortlink"
	Collection = "group"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		CreateGroup(ctx context.Context, gid string, in *user.CreateGroupRequest) error
		GetGroup(ctx context.Context, in *user.GetGroupRequest, v any) error
		UpdateGroupName(ctx context.Context, in *user.UpdateGroupRequest) error
		DeleteGroup(ctx context.Context, in *user.DeleteGroupRequest) error
		HasGid(ctx context.Context, gid, username string) bool
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

func (c *customGroupModel) HasGid(ctx context.Context, gid, username string) bool {
	filter := bson.M{
		"gid":      gid,
		"username": username,
	}

	var group Group

	err := c.conn.FindOne(ctx, &group, filter)
	if err != nil {
		return false
	}

	return true
}

func (c *customGroupModel) CreateGroup(ctx context.Context, gid string, in *user.CreateGroupRequest) error {
	filter := bson.M{
		"name": in.GetName(),
	}

	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("分组名已存在")
	}

	groupData := map[string]interface{}{
		"name":     in.Name,
		"username": in.Username,
		"gid":      gid,
		"createAt": time.Now(),
		"updateAt": time.Now(),
	}

	_, err = c.conn.InsertOne(ctx, groupData)
	if err != nil {
		return err
	}
	return nil
}

func (c *customGroupModel) GetGroup(ctx context.Context, in *user.GetGroupRequest, v any) error {
	filter := bson.M{"gid": in.GetGid()}

	err := c.conn.Find(ctx, v, filter)
	if err != nil {
		return err
	}

	return err
}

func (c *customGroupModel) UpdateGroupName(ctx context.Context, in *user.UpdateGroupRequest) error {
	filter := bson.M{"gid": in.Gid}
	update := bson.M{"$set": bson.M{"name": in.GetName()}}

	_, err := c.conn.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (c *customGroupModel) DeleteGroup(ctx context.Context, in *user.DeleteGroupRequest) error {
	filter := bson.M{"gid": in.Gid}

	_, err := c.conn.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// NewGroupModel returns a model for the mongo.
func NewGroupModel(url, db, collection string) GroupModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customGroupModel{
		defaultGroupModel: newDefaultGroupModel(conn),
	}
}
