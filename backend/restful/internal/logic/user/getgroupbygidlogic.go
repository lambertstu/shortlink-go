package user

import (
	"context"

	"restful/internal/svc"
	"restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupByGidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupByGidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByGidLogic {
	return &GetGroupByGidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupByGidLogic) GetGroupByGid(req *types.GetGroupRequest) (resp *types.GetGroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
