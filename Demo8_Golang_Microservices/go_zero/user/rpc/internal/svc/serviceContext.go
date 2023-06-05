package svc

import (
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/model"
	"Study_Golang/Demo8_Golang_Microservices/go_zero/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
	}
}
