// Code generated by goctl. DO NOT EDIT.
// Source: hello.proto

package helloclient

import (
	"context"

	"zero_shop/app/test_gateway/demo2/server/hello"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = hello.Request
	Response = hello.Response

	Hello interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultHello struct {
		cli zrpc.Client
	}
)

func NewHello(cli zrpc.Client) Hello {
	return &defaultHello{
		cli: cli,
	}
}

func (m *defaultHello) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := hello.NewHelloClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
