package logic

import (
	"context"
	"zero_shop/app/user/internal/svc"
	"zero_shop/app/user/user"
	"zero_shop/pkg"

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
	pkg.MD5(&in.Password)
	u := l.svcCtx.Uc.GetUser(context.Background(), in)
	if u == nil {
		return &user.LoginResponse{Code: 400, Msg: "用户名或密码错误"}, nil
	}
	token, _ := pkg.Sign(in.Username)
	return &user.LoginResponse{Code: 200, Msg: "success", Data: []string{
		token,
	}}, nil
}
