//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"austin-v2/app/msgpusher/internal/biz"
	"austin-v2/app/msgpusher/internal/conf"
	"austin-v2/app/msgpusher/internal/data"
	"austin-v2/app/msgpusher/internal/process"
	"austin-v2/app/msgpusher/internal/server"
	"austin-v2/app/msgpusher/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		data.DataProviderSet,
		biz.BizProviderSet,
		service.ServiceProviderSet,
		process.ProcessProviderSet,
		newApp,
	))
}
