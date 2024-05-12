package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoodsDetails struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	GoodID   string             `bson:"goodId,omitempty" json:"goodId,omitempty"`
	Data     bson.M             `bson:"data,omitempty" json:"data,omitempty"`
}

func (g *GoodsDetails) ToBytes() []byte {
	data, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	return data
}

func (g *GoodsDetails) ToString() string {
	return string(g.ToBytes())
}
