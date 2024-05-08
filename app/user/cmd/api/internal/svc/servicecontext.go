package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero_shop/app/user/cmd/api/internal/config"
	"zero_shop/app/user/cmd/rpc/user"
	"zero_shop/app/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcClient)),
	}
}
