//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/conf"
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/data"
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/server"
	"github.com/ZQCard/kratos-base-project/app/authorization/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Casbin, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
