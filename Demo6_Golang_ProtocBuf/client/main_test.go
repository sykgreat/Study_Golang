package client

import (
	"Study_Golang/Demo6_Golang_ProtocBuf/server/rpc"

	"context"
	"google.golang.org/grpc"
	"testing"
)

func Test_Main(t *testing.T) {
	dial, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
		return
	}

	client := rpc.NewServerClient(dial)

	hello, err := client.Hello(context.Background(), &rpc.Empty{})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(hello)

	//register, err := client.Register(context.Background(), &rpc.RegisterRequest{
	//	Name:     "张三",
	//	Password: "fdsjl",
	//})
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//t.Log(register)
}
