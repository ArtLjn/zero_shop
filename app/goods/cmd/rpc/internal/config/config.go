package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	RDB struct {
		Addr     string
		DB       int
		Password string
	}
}
