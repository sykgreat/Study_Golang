package main

import (
	"flag"
	"fmt"

	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/internal/config"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/internal/server"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/internal/svc"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
