package grouplogic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupByGidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupByGidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByGidLogic {
	return &GetGroupByGidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupByGidLogic) GetGroupByGid(in *user.GetGroupRequest) (*user.GetGroupResponse, error) {
	groupInfo, err := l.svcCtx.GroupModel.GetGroupByGid(l.ctx, in.GetGid())
	if err != nil {
		return nil, err
	}

	return &user.GetGroupResponse{
		Gid:      groupInfo.GID,
		Name:     groupInfo.Name,
		Username: groupInfo.Username,
	}, nil
}
