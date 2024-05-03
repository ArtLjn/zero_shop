package logic

import (
	"context"
	"zero_shop/app/payment/internal/svc"
	"zero_shop/app/payment/payment"

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

func (l *PingLogic) Ping(in *payment.Request) (*payment.Response, error) {
	// todo: add your logic here and delete this line

	return &payment.Response{}, nil
}
