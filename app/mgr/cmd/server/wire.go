//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"austin-v2/app/mgr/internal/conf"
	"austin-v2/app/mgr/internal/data"
	"austin-v2/app/mgr/internal/server"
	"austin-v2/app/mgr/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Bootstrap, *conf.Auth, *conf.Service, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		data.DataProviderSet,
		//biz.BizProviderSet,
		service.ServiceProviderSet,
		newApp,
	))
}
