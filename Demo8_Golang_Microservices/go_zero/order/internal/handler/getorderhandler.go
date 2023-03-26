package handler

import (
	"net/http"

	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/logic"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/svc"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
