package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"sync"
	"zero_shop/app/goods/cmd/rpc/good"
	"zero_shop/app/goods/cmd/rpc/internal/svc"
	"zero_shop/pkg"
)

type FindGoodPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var mu = &sync.Mutex{}

func NewFindGoodPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindGoodPageLogic {
	return &FindGoodPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindGoodPageLogic) FindGoodPage(in *good.FindGoodRequest) (*good.FindGoodResponse, error) {
	var goodList []*good.GoodData

	// TODO 思考：如果查询每页数量与实际缓存页码中数量不一致问题
	// 进行redis数据查询
	goodCache, err := l.svcCtx.Rdb.Hget(pkg.GoodsKey, strconv.Itoa(int(in.Page)))
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// 如果redis中没有数据，则进行数据库查询，并将数据存入redis中
	if len(goodCache) == 0 {
		mu.Lock()
		defer mu.Unlock()
		data, qe := l.svcCtx.Gc.FindGoodPage(l.ctx, in.Page, in.Size)
		if qe != nil {
			return nil, qe
		}
		go func() {
			if len(data) != 0 {
				by, _ := json.Marshal(&data)
				ex := l.svcCtx.Rdb.Hset(pkg.GoodsKey, strconv.Itoa(int(in.Page)), string(by))
				if ex != nil {
					logx.Error(ex)
				}
			}
		}()
		l.svcCtx.Tool.CopyProperties(data, &goodList)
	}

	l.svcCtx.Tool.ParseList(goodCache, &goodList)
	return &good.FindGoodResponse{
		Code: 200,
		Msg:  "success",
		Data: goodList,
	}, nil
}
