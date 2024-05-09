package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero_shop/app/goods/cmd/api/internal/config"
	"zero_shop/app/goods/cmd/rpc/good"
	"zero_shop/app/goods/cmd/rpc/goodclient"
	"zero_shop/app/user/cmd/rpc/user"
	"zero_shop/app/user/cmd/rpc/userclient"
	"zero_shop/pkg"
)

type ServiceContext struct {
	Config  config.Config
	GoodRpc good.GoodClient
	UserRpc user.UserClient
	Tool    pkg.HuToolUtils
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		GoodRpc: goodclient.NewGood(zrpc.MustNewClient(c.GoodRpcClient)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcClient)),
		Tool:    pkg.NewBeanUtils(),
	}
}
