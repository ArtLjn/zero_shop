package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"zero_shop/app/goods/cmd/rpc/good"
	"zero_shop/app/goods/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodDetailsLogic {
	return &GetGoodDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodDetailsLogic) GetGoodDetails(in *good.GetGoodDetailsRequest) (*good.GetGoodDetailsResponse, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.Mongo.FindOneGoodDeatilsByCond(l.ctx, bson.M{"goodId": in.GetGoodId()})
	if err != nil {
		return nil, err
	}
	return &good.GetGoodDetailsResponse{
		Code: 200,
		Msg:  "success",
		Data: data.ToString(),
	}, nil
}
