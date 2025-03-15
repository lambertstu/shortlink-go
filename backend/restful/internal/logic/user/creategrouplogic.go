package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupRequest) (resp *types.CreateGroupResponse, err error) {
	rpcResp, err := l.svcCtx.GroupRpcService.CreateGroup(l.ctx, &user.CreateGroupRequest{
		Name:     req.Name,
		Username: req.Username,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("创建分组失败")
	}
	return &types.CreateGroupResponse{
		Gid:      rpcResp.Gid,
		Name:     rpcResp.Name,
		Username: rpcResp.Username,
	}, nil
}
