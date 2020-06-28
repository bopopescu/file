package service

import (
	"jupiter-demo/internal/app/model"
	"jupiter-demo/internal/app/service/user"
	"jupiter-demo/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)
//Init instantiate the service
func Init()  {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}