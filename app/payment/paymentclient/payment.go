// Code generated by goctl. DO NOT EDIT.
// Source: payment.proto

package paymentclient

import (
	"context"
	"zero_shop/app/payment/payment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = payment.Request
	Response = payment.Response

	Payment interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultPayment struct {
		cli zrpc.Client
	}
)

func NewPayment(cli zrpc.Client) Payment {
	return &defaultPayment{
		cli: cli,
	}
}

func (m *defaultPayment) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
