package main

import (
	helloGrpc "Study_Golang/Demo6_Golang_ProtocBuf/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	helloGrpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(_ context.Context, req *helloGrpc.Req) (res *helloGrpc.Res, err error) {
	log.Printf("SayHi() is called with: %v\n", req.GetMessage())
	return &helloGrpc.Res{Message: "Hello, server"}, nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	l, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	helloGrpc.RegisterHelloGRPCServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
