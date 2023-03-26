package svc

import (
	"Study_Golang/Demo8_Golang_Microservices/go_zero/order/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
