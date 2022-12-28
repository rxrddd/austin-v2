//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/conf"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/data"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/server"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, *conf.Service, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
