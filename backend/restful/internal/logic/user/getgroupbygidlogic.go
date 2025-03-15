package user

import (
	"context"
	"errors"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

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
	userInfo, err := l.svcCtx.GroupRpcService.GetGroupByGid(l.ctx, &user.GetGroupRequest{
		Gid: req.Gid,
	})

	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("获取分组信息失败")
	}
	return &types.GetGroupResponse{
		Gid:      userInfo.Gid,
		Name:     userInfo.Name,
		Username: userInfo.Username,
	}, nil
}
