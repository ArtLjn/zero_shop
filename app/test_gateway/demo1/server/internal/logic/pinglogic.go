package logic

import (
	"context"
	"zero_shop/app/test_gateway/demo1/server/hello"
	"zero_shop/app/test_gateway/demo1/server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *PingLogic) Ping(in *hello.Request) (*hello.Response, error) {
	return &hello.Response{
		Msg: "pong",
	}, nil
}
