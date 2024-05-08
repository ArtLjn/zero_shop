package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero_shop/app/user/cmd/api/internal/logic"
	"zero_shop/app/user/cmd/api/internal/svc"
	"zero_shop/app/user/cmd/rpc/user"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req user.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
