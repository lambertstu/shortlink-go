package grouplogic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupLogic {
	return &GetGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupLogic) GetGroup(in *user.GetGroupRequest) (*user.GetGroupResponse, error) {
	var groupList []*user.GroupData
	err := l.svcCtx.GroupModel.GetGroup(l.ctx, in, groupList)
	if err != nil {
		return nil, err
	}

	return &user.GetGroupResponse{
		Gid:  in.GetGid(),
		Data: groupList,
	}, nil
}
