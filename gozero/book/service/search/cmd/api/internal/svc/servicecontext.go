package svc

import (
	"StandardProject/gozero/book/service/search/cmd/api/internal/config"
	"StandardProject/gozero/book/service/search/cmd/api/internal/middleware"
	"StandardProject/gozero/book/service/search/cmd/api/internal/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
