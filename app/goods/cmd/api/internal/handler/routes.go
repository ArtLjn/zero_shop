// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"zero_shop/pkg"
	"zero_shop/pkg/middleware"

	"zero_shop/app/goods/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodPost,
				Path:   "/api/goods/createGood",
				Handler: middleware.AuthHandler(middleware.CheckRoleHandler(
					CreateGoodHandler(serverCtx),
					serverCtx.UserRpc,
					pkg.SellerRole,
				),
				),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/goods/getGoods",
				Handler: FindGoodHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/goods/createGoodDetails",
				Handler: CreateGoodDetailsHandler(serverCtx),
			},
		},
	)
}
