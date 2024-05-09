package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"zero_shop/app/goods/model"

	"zero_shop/app/goods/cmd/rpc/good"
	"zero_shop/app/goods/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodLogic {
	return &CreateGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateGood ... 创建商品
func (l *CreateGoodLogic) CreateGood(in *good.CreateGoodRequest) (*good.CreateGoodResponse, error) {
	var gd model.Goods
	l.svcCtx.Tool.CopyProperties(in, &gd)
	if l.svcCtx.Tool.IsEmpty(gd, &model.Goods{}) {
		return nil, fmt.Errorf("goods is empty")
	}
	gd.GoodId = uuid.New().String()[:8]
	_, err := l.svcCtx.Gc.Insert(context.Background(), &gd)
	if err != nil {
		return nil, err
	}
	return &good.CreateGoodResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
