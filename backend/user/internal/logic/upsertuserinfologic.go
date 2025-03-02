package logic

import (
	"context"
	model "user/mongo/user"
	"user/pkg/constant"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertUserInfoLogic {
	return &UpsertUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpsertUserInfoLogic) UpsertUserInfo(in *user.UpsertUserInfoRequest) (*user.UpsertUserInfoResponse, error) {
	err := l.svcCtx.UserModel.UpdateUserInfo(l.ctx, &model.User{
		Username:   in.Username,
		Password:   in.Password,
		RealName:   in.RealName,
		Phone:      in.Phone,
		Mail:       in.Email,
		DeleteFlag: constant.DELETE_FLAG,
	})
	if err != nil {
		return &user.UpsertUserInfoResponse{
			Success: false,
		}, err
	}
	return &user.UpsertUserInfoResponse{
		Success: true,
	}, nil
}
