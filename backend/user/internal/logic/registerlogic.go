package logic

import (
	"context"
	//"github.com/Lambert1215/shortlink-go/backend/pkg/constant"
	model "user/mongo/user"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l.svcCtx.UserModel.Register(context.Background(), &model.User{
		Username:   in.Username,
		Password:   in.Password,
		RealName:   in.RealName,
		Phone:      in.Phone,
		Mail:       in.Email,
		DeleteFlag: 0,
	})
	return &user.RegisterResponse{}, nil
}
