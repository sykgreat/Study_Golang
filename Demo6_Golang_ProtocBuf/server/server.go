package server

import (
	"Study_Golang/Demo6_Golang_ProtocBuf/server/rpc"

	"context"
	"fmt"
)

type Server struct{}

func (s *Server) Hello(context.Context, *rpc.Empty) (*rpc.HelloResponse, error) {
	return &rpc.HelloResponse{Message: "Hello client ..."}, nil
}

func (s *Server) Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	return &rpc.RegisterResponse{
		Uid: fmt.Sprintf("%s_%s", request.GetName(), request.GetPassword()),
	}, nil
}
