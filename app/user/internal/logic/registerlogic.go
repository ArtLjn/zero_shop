package logic

import (
	"context"
	"zero_shop/app/user/user"
	"zero_shop/pkg"

	"github.com/zeromicro/go-zero/core/logx"
	"zero_shop/app/user/internal/svc"
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
	// validate fields is not empty
	if err := pkg.CheckFieldsNoEmpty(in); err != nil {
		return &user.RegisterResponse{
			Code: 400,
			Msg:  err.Error(),
		}, err
	}
	// check register email code
	cacheCode, _ := l.svcCtx.Rdb.Get(context.Background(), in.Email).Result()
	if cacheCode != in.EmailCode {
		return &user.RegisterResponse{
			Code: 400,
			Msg:  "register code is wrong",
		}, nil
	}

	return &user.RegisterResponse{}, nil
}
