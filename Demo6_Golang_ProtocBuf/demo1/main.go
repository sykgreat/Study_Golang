package main

import (
	"Study_Golang/Demo6_Golang_ProtocBuf/demo1/service"
	"github.com/golang/protobuf/proto"
)

func protoc() {
	user1 := &service.User{
		Name: "Tom",
		Age:  18,
	}

	// 序列化
	marshal, err := proto.Marshal(user1)
	if err != nil {
		panic(err)
	}

	// 反序列化
	user2 := &service.User{}
	if err = proto.Unmarshal(marshal, user2); err != nil {
		panic(err)
	}
}

func main() {
	protoc()
}
