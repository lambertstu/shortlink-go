package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ShortlinkModel = (*customShortlinkModel)(nil)

type (
	// ShortlinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortlinkModel.
	ShortlinkModel interface {
		shortlinkModel
	}

	customShortlinkModel struct {
		*defaultShortlinkModel
	}
)

// NewShortlinkModel returns a model for the mongo.
func NewShortlinkModel(url, db, collection string) ShortlinkModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customShortlinkModel{
		defaultShortlinkModel: newDefaultShortlinkModel(conn),
	}
}
