package logic

import (
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/internal/svc"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/types/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserResponse{
		Id:     in.Id,
		Name:   "test",
		Gender: "男",
	}, nil
}
