package main

import (
	"flag"
	"fmt"
	"zero_shop/pkg"

	"zero_shop/app/goods/cmd/api/internal/config"
	"zero_shop/app/goods/cmd/api/internal/handler"
	"zero_shop/app/goods/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", pkg.GoodsHttpConfigPath, "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
