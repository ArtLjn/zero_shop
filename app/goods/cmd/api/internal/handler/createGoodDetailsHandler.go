package handler

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"zero_shop/app/goods/cmd/api/internal/logic"
	"zero_shop/app/goods/cmd/api/internal/svc"
	"zero_shop/app/goods/cmd/rpc/good"
)

func CreateGoodDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req good.CreateGoodDetailsRequest
		str, err := io.ReadAll(r.Body)
		if err = json.Unmarshal(str, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewGoodLogic(r.Context(), svcCtx)
		resp, err := l.CreateGoodDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
