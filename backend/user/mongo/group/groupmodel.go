package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"user/pb/user"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		CreateGroup(ctx context.Context, gid string, in *user.CreateGroupRequest) error
		GetGroup(ctx context.Context, in *user.GetGroupRequest) (*user.GetGroupResponse, error)
		UpdateGroup(ctx context.Context, in *user.UpdateGroupRequest) error
		DeleteGroup(ctx context.Context, in *user.DeleteGroupRequest) error
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

func (c *customGroupModel) CreateGroup(ctx context.Context, gid string, in *user.CreateGroupRequest) error {
	groupData := map[string]interface{}{
		"name":     in.Name,
		"username": in.Username,
		"gid":      gid,
	}

	_, err := c.conn.InsertOne(ctx, groupData)
	if err != nil {
		return err
	}
	return nil
}

func (c *customGroupModel) GetGroup(ctx context.Context, in *user.GetGroupRequest) (*user.GetGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *customGroupModel) UpdateGroup(ctx context.Context, in *user.UpdateGroupRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c *customGroupModel) DeleteGroup(ctx context.Context, in *user.DeleteGroupRequest) error {
	//TODO implement me
	panic("implement me")
}

// NewGroupModel returns a model for the mongo.
func NewGroupModel(url, db, collection string) GroupModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customGroupModel{
		defaultGroupModel: newDefaultGroupModel(conn),
	}
}
