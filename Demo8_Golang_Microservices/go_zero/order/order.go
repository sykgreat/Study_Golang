package main

import (
	"flag"
	"fmt"

	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/config"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/handler"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

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
