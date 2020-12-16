// wire.go
//+build wireinject

package main

import (
	"github.com/B8kingS0ga/Go-000/tree/main/Week04/internal/api/dao"
	"github.com/B8kingS0ga/Go-000/tree/main/Week04/internal/api/handler"
	"github.com/B8kingS0ga/Go-000/tree/main/Week04/internal/api/service"
	"github.com/google/wire"
)

func InitializeServer() handler.HelloHandle {
	wire.Build(handler.NewHandle, service.NewService, dao.NewDao)
	return handler.HelloHandle{}
}
