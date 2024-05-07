package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Email struct {
		Username string
		Password string
		Port     int
		Host     string
	}
	RDB struct {
		Addr     string
		DB       int
		Password string
	}
}
