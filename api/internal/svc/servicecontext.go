package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-chat/api/internal/config"
	"zero-chat/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(c.Auth.AccessSecret).Handle,
	}
}
