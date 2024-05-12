package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	RDB     redis.RedisConf
	MongoDB struct {
		Url        string
		Db         string
		Collection string
	}
}
