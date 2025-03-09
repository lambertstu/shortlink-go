package userlogic

import (
	"context"
	"errors"
	model "user/mongo/user"
	"user/pkg/constant"
	"user/pkg/errorcode"

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
	updateRes, err := l.svcCtx.UserModel.UpdateUserInfo(l.ctx, &model.User{
		Username:   in.Username,
		Password:   in.Password,
		RealName:   in.RealName,
		Phone:      in.Phone,
		Mail:       in.Email,
		DeleteFlag: constant.ENABLE_FLAG,
	})
	if err != nil || updateRes.MatchedCount == 0 || updateRes.ModifiedCount == 0 {
		if updateRes.MatchedCount == 0 {
			return &user.UpsertUserInfoResponse{
				Success: false,
			}, errors.New(errorcode.UserNotExist.Message())
		} else if updateRes.ModifiedCount == 0 {
			return &user.UpsertUserInfoResponse{
				Success: false,
			}, errors.New("修改信息失败")
		}

		return &user.UpsertUserInfoResponse{
			Success: false,
		}, err
	}

	return &user.UpsertUserInfoResponse{
		Success: true,
	}, nil
}
