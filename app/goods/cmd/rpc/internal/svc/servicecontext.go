package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	model2 "zero_shop/app/goods/cache"
	"zero_shop/app/goods/cmd/rpc/internal/config"
	"zero_shop/app/goods/model"
	"zero_shop/pkg"
)

type ServiceContext struct {
	Config config.Config
	Tool   pkg.HuToolUtils
	Gc     model.GoodsModel
	Rdb    *redis.Redis
	Mongo  model2.GoodsDetailsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Tool:   pkg.NewBeanUtils(),
		Gc:     model.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
		Rdb:    redis.MustNewRedis(c.RDB),
		Mongo: model2.NewGoodsDetailsModel(c.MongoDB.Url, c.MongoDB.Db, c.MongoDB.Collection, cache.CacheConf{
			cache.NodeConf{
				RedisConf: c.RDB,
				Weight:    100,
			},
		}),
	}
}
