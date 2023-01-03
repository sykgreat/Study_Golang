package main

import (
	helloGrpc "Study_Golang/Demo6_Golang_ProtocBuf/pb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	dial, err := grpc.Dial(":8008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v\n", err)
	}
	defer func(dial *grpc.ClientConn) {
		err := dial.Close()
		if err != nil {
			log.Fatalf("failed to close dial: %v\n", err)
		}
	}(dial)
	client := helloGrpc.NewHelloGRPCClient(dial)

	hi, err := client.SayHi(context.Background(), &helloGrpc.Req{Message: "Hello, client"})
	if err != nil {
		log.Fatalf("failed to SayHi: %v\n", err)
	}
	log.Println(hi.GetMessage())
}
