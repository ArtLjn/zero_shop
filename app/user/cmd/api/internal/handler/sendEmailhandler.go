package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero_shop/app/user/cmd/api/internal/logic"
	"zero_shop/app/user/cmd/api/internal/svc"
	"zero_shop/app/user/cmd/rpc/user"
)

func SendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req user.SendEmailRequest
		req.Email = r.URL.Query().Get("email")
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.SendEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
