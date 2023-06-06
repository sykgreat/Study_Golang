package logic

import (
	"context"
	"errors"

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
	login, err := l.svcCtx.UserModel.Login(l.ctx, in.UserName, in.Password)
	if err != nil {
		return nil, err
	}
	if login == nil {
		return nil, errors.New("用户名或密码错误")
	}
	return &pb.LoginResp{}, nil
}
