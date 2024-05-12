package logic

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	model "zero_shop/app/goods/cache"
	"zero_shop/app/goods/cmd/rpc/good"
	"zero_shop/app/goods/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodDetailsLogic {
	return &CreateGoodDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodDetailsLogic) CreateGoodDetails(in *good.CreateGoodDetailsRequest) (*good.CreateGoodDetailsResponse, error) {
	var data bson.M
	if je := json.Unmarshal([]byte(in.Data), &data); je != nil {
		return nil, je
	}
	err := l.svcCtx.Mongo.Insert(context.Background(), &model.GoodsDetails{
		ID:       primitive.NewObjectID(),
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
		GoodID:   in.GetGoodId(),
		Data:     data,
	})
	if err != nil {
		panic(err)
	}
	return &good.CreateGoodDetailsResponse{}, nil
}
