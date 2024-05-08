package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"zero_shop/app/order/internal/svc"
	"zero_shop/app/order/order"
	"zero_shop/app/user/cmd/rpc/user"
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

func (l *PingLogic) Ping(in *order.Request) (*order.Response, error) {
	// todo: add your logic here and delete this line
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc",
		},
	})
	client := user.NewUserClient(conn.Conn())
	resp, err := client.Login(context.Background(), &user.LoginRequest{
		Username: "123",
		Password: "123",
	})
	fmt.Println(resp, err)
	return &order.Response{Pong: in.Ping}, nil
}
