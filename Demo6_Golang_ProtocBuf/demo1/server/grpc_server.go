package main

import (
	"Study_Golang/Demo6_Golang_ProtocBuf/demo1/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//tlsFile, err := credentials.NewClientTLSFromFile("Demo6_Golang_ProtocBuf/demo1/cert/cert.pem", "Demo6_Golang_ProtocBuf/demo1/cert/private.key")
	//if err != nil {
	//	log.Fatalln("failed to create tls: ", err)
	//}
	//
	//rpcServer := grpc.NewServer(grpc.Creds(tlsFile))
	rpcServer := grpc.NewServer()

	service.RegisterProductServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	err = rpcServer.Serve(listen)
	if err != nil {
		log.Fatalln("failed to serve: ", err)
	}

	fmt.Println("Server is running...")
}
