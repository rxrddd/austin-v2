// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"austin-v2/app/msgpusher-manager/internal/biz"
	"austin-v2/app/msgpusher-manager/internal/conf"
	"austin-v2/app/msgpusher-manager/internal/data"
	"austin-v2/app/msgpusher-manager/internal/server"
	"austin-v2/app/msgpusher-manager/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	iMessagingClient := data.NewMq(confData, logger)
	db := data.NewMysqlCmd(confData, logger)
	cmdable := data.NewRedisCmd(confData, logger)
	client := data.NewMongoDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger, iMessagingClient, db, cmdable, client)
	if err != nil {
		return nil, nil, err
	}
	iMsgRecordRepo := data.NewMsgRecordRepo(dataData, logger)
	msgRecordUseCase := biz.NewMsgRecordUseCase(iMsgRecordRepo, logger)
	iMessageTemplateRepo := data.NewMessageTemplateRepo(dataData, logger)
	messageTemplateUseCase := biz.NewMessageTemplateUseCase(iMessageTemplateRepo, logger)
	iSmsRecordRepo := data.NewSmsRecordRepo(dataData, logger)
	smsRecordUseCase := biz.NewSmsRecordUseCase(iSmsRecordRepo, logger)
	iSendAccountRepo := data.NewSendAccountRepo(dataData, logger)
	sendAccountUseCase := biz.NewSendAccountUseCase(iSendAccountRepo, logger)
	msgPusherManagerService := service.NewMsgPusherManagerService(msgRecordUseCase, messageTemplateUseCase, smsRecordUseCase, sendAccountUseCase)
	grpcServer := server.NewGRPCServer(confServer, msgPusherManagerService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}