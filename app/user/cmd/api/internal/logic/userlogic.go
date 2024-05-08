package logic

import (
	"context"
	"zero_shop/app/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"zero_shop/app/user/cmd/api/internal/svc"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Login(req *user.LoginRequest) (*user.LoginResponse, error) {
	resp, err := l.svcCtx.UserRpc.Login(l.ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *UserLogic) SendEmail(req *user.SendEmailRequest) (*user.SendEmailResponse, error) {
	resp, err := l.svcCtx.UserRpc.SendEmail(l.ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *UserLogic) Register(req *user.RegisterRequest) (*user.RegisterResponse, error) {
	resp, err := l.svcCtx.UserRpc.Register(l.ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
