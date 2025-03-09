package grouplogic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(in *user.DeleteGroupRequest) (*user.DeleteGroupResponse, error) {
	err := l.svcCtx.GroupModel.DeleteGroup(l.ctx, in)
	if err != nil {
		return &user.DeleteGroupResponse{
			Success: false,
		}, err
	}
	return &user.DeleteGroupResponse{
		Success: true,
	}, nil
}
