package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupLogic) DeleteGroup(req *types.DeleteGroupRequest) (resp *types.NilResponse, err error) {
	rpcResp, _ := l.svcCtx.GroupRpcService.DeleteGroup(l.ctx, &user.DeleteGroupRequest{
		Gid:      req.Gid,
		Username: req.Username,
	})
	if !rpcResp.Success {
		return nil, errors.New("删除分组失败")
	}
	return nil, nil
}
