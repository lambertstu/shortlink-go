package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/lambertstu/shortlink-user-rpc/client/user"

	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupByUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupByUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupByUserNameLogic {
	return &GetGroupByUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupByUserNameLogic) GetGroupByUserName(req *types.GetGroupRequest) (resp *types.GetGroupResponse, err error) {
	groups, err := l.svcCtx.GroupRpcService.GetGroupByUsername(l.ctx, &user.GetGroupRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, errors.New("获取分组失败")
	}

	fmt.Println(groups)
	var groupList []types.GroupData
	for _, group := range groups.Data {
		groupList = append(groupList, types.GroupData{
			Gid:      group.Gid,
			Name:     group.Name,
			Username: group.Username,
		})
	}

	return &types.GetGroupResponse{
		Data: groupList,
	}, nil
}
