package logic

import (
	"context"
	"fmt"
	"zero_shop/app/user/user"

	"zero_shop/app/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.Username, in.Password)
	return &user.LoginResponse{Code: 1, Msg: "success", Data: []string{in.Username, in.Password}}, nil
}
