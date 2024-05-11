package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoodsDetails struct {
	ID       primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt time.Time              `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time              `bson:"createAt,omitempty" json:"createAt,omitempty"`
	Data     map[string]interface{} `bson:"data,omitempty" json:"data,omitempty"`
}
