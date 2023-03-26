package logic

import (
	"context"

	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/svc"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderRequest) (resp *types.OrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
