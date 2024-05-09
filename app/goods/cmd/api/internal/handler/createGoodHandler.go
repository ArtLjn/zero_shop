package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero_shop/app/goods/cmd/api/internal/logic"
	"zero_shop/app/goods/cmd/api/internal/svc"
	"zero_shop/app/goods/cmd/api/internal/types"
)

func CreateGoodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateGoodRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewGoodLogic(r.Context(), svcCtx)
		resp, err := l.CreateGood(&req, r.Context().Value("userId").(string))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
