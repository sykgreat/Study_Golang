// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/internal/logic"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/internal/svc"
	user2 "Study_Golang/Demo8_Golang_Microservices/go_zero/user/types/user"
	"context"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user2.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user2.IdRequest) (*user2.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}