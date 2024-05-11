package logic

import (
	"context"
	"zero_shop/app/goods/cmd/api/internal/types"
	"zero_shop/app/goods/cmd/rpc/good"

	"github.com/zeromicro/go-zero/core/logx"
	"zero_shop/app/goods/cmd/api/internal/svc"
)

type GoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodLogic {
	return &GoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (g *GoodLogic) CreateGood(req *types.CreateGoodRequest, sellerId string) (*good.CreateGoodResponse, error) {
	var goodReq *good.CreateGoodRequest
	g.svcCtx.Tool.CopyProperties(req, &goodReq)
	goodReq.SellerId = sellerId
	resp, err := g.svcCtx.GoodRpc.CreateGood(g.ctx, goodReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GoodLogic) FindGoodPage(req *good.FindGoodRequest) (*good.FindGoodResponse, error) {
	return g.svcCtx.GoodRpc.FindGoodPage(g.ctx, req)
}

func (g *GoodLogic) CreateGoodDetails(req *good.CreateGoodDetailsRequest) (*good.CreateGoodDetailsResponse, error) {
	return g.svcCtx.GoodRpc.CreateGoodDetails(g.ctx, req)
}
