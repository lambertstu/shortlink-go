package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.UpdateGroupRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.GroupRpcService.UpdateGroup(l.ctx, &user.UpdateGroupRequest{
		Gid:      req.Gid,
		Name:     req.Name,
		Username: req.Username,
	})
	if !rpcResp.Success {
		return nil, errors.New("更新分组信息失败")
	}
	return nil, nil
}
