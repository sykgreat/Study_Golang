package main

//import "C"
import (
	"Study_Golang/Demo6_Golang_ProtocBuf/demo1/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//
	//tlsFile, err := credentials.NewClientTLSFromFile("Demo6_Golang_ProtocBuf/demo1/cert/cert.pem", "127.0.0.1")
	//if err != nil {
	//	log.Fatalln("failed to create tls: ", err)
	//}
	//
	//conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(tlsFile))
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("failed to dial: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln("failed to close: ", err)
		}
	}(conn)

	client := service.NewProductServiceClient(conn)

	request := service.ProductRequest{
		ProductId: 123,
	}

	productResponse, err := client.GetProductStock(context.Background(), &request)
	if err != nil {
		log.Fatalln("failed to get product stock: ", err)
	}
	fmt.Println("Product stock: ", productResponse.ProductStock)
}
