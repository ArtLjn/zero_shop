package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ GoodsDetailsModel = (*customGoodsDetailsModel)(nil)

type (
	// GoodsDetailsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsDetailsModel.
	GoodsDetailsModel interface {
		goodsDetailsModel
	}

	customGoodsDetailsModel struct {
		*defaultGoodsDetailsModel
	}
)

// NewGoodsDetailsModel returns a model for the mongo.
func NewGoodsDetailsModel(url, db, collection string, c cache.CacheConf) GoodsDetailsModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customGoodsDetailsModel{
		defaultGoodsDetailsModel: newDefaultGoodsDetailsModel(conn),
	}
}
