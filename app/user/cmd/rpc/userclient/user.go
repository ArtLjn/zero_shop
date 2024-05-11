// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"zero_shop/app/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserMsgRequest  = user.GetUserMsgRequest
	GetUserMsgResponse = user.GetUserMsgResponse
	LoginRequest       = user.LoginRequest
	LoginResponse      = user.LoginResponse
	RegisterRequest    = user.RegisterRequest
	RegisterResponse   = user.RegisterResponse
	SendEmailRequest   = user.SendEmailRequest
	SendEmailResponse  = user.SendEmailResponse

	User interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
		GetUserMsg(ctx context.Context, in *GetUserMsgRequest, opts ...grpc.CallOption) (*GetUserMsgResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SendEmail(ctx, in, opts...)
}

func (m *defaultUser) GetUserMsg(ctx context.Context, in *GetUserMsgRequest, opts ...grpc.CallOption) (*GetUserMsgResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserMsg(ctx, in, opts...)
}