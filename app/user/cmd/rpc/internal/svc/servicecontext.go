package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero_shop/app/user/cmd/rpc/internal/config"
	"zero_shop/app/user/model"
	"zero_shop/pkg"
)

type ServiceContext struct {
	Config config.Config
	Uc     model.UserModel
	Rdb    *redis.Client
	Tool   pkg.HuToolUtils
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Uc:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		Rdb:    NewRDB(c),
		Tool:   pkg.NewBeanUtils(),
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
