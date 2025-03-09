package grouplogic

import (
	"context"
	"user/internal/svc"
	"user/pb/user"
	"user/pkg/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *user.CreateGroupRequest) (*user.CreateGroupResponse, error) {
	var (
		gid string
		err error
	)

	for {
		gid, err = tool.GenerateRandomSequence()
		if err != nil {
			return nil, err
		}

		hasGid := l.svcCtx.GroupModel.HasGid(l.ctx, gid, in.GetUsername())
		if hasGid {
			break
		}
	}

	err = l.svcCtx.GroupModel.CreateGroup(l.ctx, gid, in)
	if err != nil {
		return nil, err
	}

	return &user.CreateGroupResponse{
		Gid:      gid,
		Name:     in.GetName(),
		Username: in.GetUsername(),
		Success:  true,
	}, nil
}
