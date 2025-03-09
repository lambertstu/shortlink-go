package grouplogic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupLogic) UpdateGroup(in *user.UpdateGroupRequest) (*user.UpdateGroupResponse, error) {
	err := l.svcCtx.GroupModel.UpdateGroupName(l.ctx, in)

	if err != nil {
		return &user.UpdateGroupResponse{
			Success: false,
		}, err
	}

	return &user.UpdateGroupResponse{
		Success: true,
	}, nil
}
