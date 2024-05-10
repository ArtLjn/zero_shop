package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
	"zero_shop/app/goods/cmd/api/internal/logic"
	"zero_shop/app/goods/cmd/api/internal/svc"
	"zero_shop/app/goods/cmd/rpc/good"
)

func FindGoodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req good.FindGoodRequest
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		size, _ := strconv.Atoi(r.URL.Query().Get("size"))
		req.Page = int64(page)
		req.Size = int64(size)
		l := logic.NewGoodLogic(r.Context(), svcCtx)
		resp, err := l.FindGoodPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
