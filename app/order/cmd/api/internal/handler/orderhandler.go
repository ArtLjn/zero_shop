package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero_shop/app/order/cmd/api/internal/logic"
	"zero_shop/app/order/cmd/api/internal/svc"
	"zero_shop/app/order/cmd/api/internal/types"
)

func OrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderLogic(r.Context(), svcCtx)
		resp, err := l.Order(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
