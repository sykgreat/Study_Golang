package logic

import (
	"context"

	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/internal/svc"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	return &pb.LoginResp{}, nil
}
