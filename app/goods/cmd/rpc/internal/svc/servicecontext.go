package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero_shop/app/goods/cmd/rpc/internal/config"
	"zero_shop/app/goods/model"
	"zero_shop/pkg"
)

type ServiceContext struct {
	Config config.Config
	Tool   pkg.HuToolUtils
	Gc     model.GoodsModel
	Rdb    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Tool:   pkg.NewBeanUtils(),
		Gc:     model.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
		Rdb:    NewRDB(c),
	}
}

func NewRDB(c config.Config) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     c.RDB.Addr,
			DB:       c.RDB.DB,
			Password: c.RDB.Password,
		},
	)
}
