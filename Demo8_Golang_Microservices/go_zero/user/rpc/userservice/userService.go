// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoReq  = pb.GetUserInfoReq
	GetUserInfoResp = pb.GetUserInfoResp
	LoginReq        = pb.LoginReq
	LoginResp       = pb.LoginResp

	UserService interface {
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
