package server

import (
	"Study_Golang/Demo6_Golang_ProtocBuf/server/rpc"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		t.Fatal(err)
		return
	}

	// 服务端多拦截器
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LogUnaryServerInterceptor(), AuthUnaryServerInterceptor())))

	rpc.RegisterServerServer(server, &Server{})

	t.Log("server start")
	if err := server.Serve(listen); err != nil {
		t.Fatal("启动 grpc 服务失败", err)
		return
	}
}

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Info().Str("method", info.FullMethod).Msg("LogUnaryServerInterceptor")
		return handler(ctx, req)
	}
}

func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Info().Str("method", info.FullMethod).Msg("AuthUnaryServerInterceptor")
		return handler(ctx, req)
	}
}
