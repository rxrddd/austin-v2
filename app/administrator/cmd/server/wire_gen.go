// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ZQCard/kratos-base-project/app/administrator/internal/biz"
	"github.com/ZQCard/kratos-base-project/app/administrator/internal/conf"
	"github.com/ZQCard/kratos-base-project/app/administrator/internal/data"
	data2 "github.com/ZQCard/kratos-base-project/app/administrator/internal/data"
	"github.com/ZQCard/kratos-base-project/app/administrator/internal/server"
	"github.com/ZQCard/kratos-base-project/app/administrator/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	db := data2.NewMysqlCmd(confData, logger)
	cmdable := data2.NewRedisCmd(confData, logger)
	dataData, cleanup, err := data2.NewData(db, cmdable, logger)
	if err != nil {
		return nil, nil, err
	}
	ServiceNameRepo := data.NewAdministratorRepo(dataData, logger)
	ServiceNameUseCase := biz.NewAdministratorUseCase(ServiceNameRepo, logger)
	ServiceNameService := service.NewAdministratorService(ServiceNameUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, ServiceNameService, logger)
	registrar := data2.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}