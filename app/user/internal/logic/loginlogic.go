package logic

import (
	"context"
	"zero_shop/app/user/model"
	"zero_shop/app/user/user"

	"zero_shop/app/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	uc model.UserModel
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO QQ登录api

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	u := l.uc.GetUser(context.Background(), in)
	if u == nil {
		return &user.LoginResponse{Code: 0, Msg: "用户名或密码错误"}, nil
	}
	return &user.LoginResponse{Code: 1, Msg: "success", Data: []string{in.Username, in.Password}}, nil
}
