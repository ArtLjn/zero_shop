// Code generated by goctl. DO NOT EDIT.
// Source: hello.proto

package server

import (
	"context"

	"zero_shop/app/test_gateway/demo2/server/hello"
	"zero_shop/app/test_gateway/demo2/server/internal/logic"
	"zero_shop/app/test_gateway/demo2/server/internal/svc"
)

type HelloServer struct {
	svcCtx *svc.ServiceContext
	hello.UnimplementedHelloServer
}

func NewHelloServer(svcCtx *svc.ServiceContext) *HelloServer {
	return &HelloServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloServer) Ping(ctx context.Context, in *hello.Request) (*hello.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
