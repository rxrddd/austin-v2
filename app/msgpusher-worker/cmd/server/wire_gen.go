// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/conf"
	"austin-v2/app/msgpusher-worker/internal/data"
	"austin-v2/app/msgpusher-worker/internal/sender"
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/server"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/app/msgpusher-worker/internal/service/deduplication"
	"austin-v2/app/msgpusher-worker/internal/service/limiter"
	"austin-v2/app/msgpusher-worker/internal/service/srv"
	"austin-v2/pkg/utils/mqHelper"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	broker := data.NewBroker(confData, logger)
	taskExecutor := sender.NewTaskExecutor()
	cmdable := data.NewRedisCmd(confData, logger)
	smsHandler := handler.NewSmsHandler(logger, cmdable)
	mqHelperMqHelper := mqHelper.NewMqHelper(broker)
	db := data.NewMysqlCmd(confData, logger)
	dataData, cleanup, err := data.NewData(confData, logger, broker, mqHelperMqHelper, cmdable, db)
	if err != nil {
		return nil, nil, err
	}
	iSendAccountRepo := data.NewSendAccountRepo(dataData, logger)
	sendAccountUseCase := biz.NewSendAccountUseCase(iSendAccountRepo, logger)
	emailHandler := handler.NewEmailHandler(logger, cmdable, sendAccountUseCase)
	handleManager := sender.NewHandleManager(smsHandler, emailHandler)
	discardMessageService := srv.NewDiscardMessageService(logger, cmdable)
	shieldService := srv.NewShieldService(logger, cmdable)
	iMessageTemplateRepo := data.NewMessageTemplateRepo(dataData, logger)
	messageTemplateUseCase := biz.NewMessageTemplateUseCase(iMessageTemplateRepo, logger)
	simpleLimitService := limit.NewSimpleLimitService(logger, cmdable)
	slideWindowLimitService := limit.NewSlideWindowLimitService(logger, confData, cmdable)
	limiterManager := limit.NewLimiterManager(simpleLimitService, slideWindowLimitService)
	frequencyDeduplicationService := deduplication.NewFrequencyDeduplicationService(limiterManager)
	contentDeduplicationService := deduplication.NewContentDeduplicationService(limiterManager)
	deduplicationManager := deduplication.NewDeduplicationManager(frequencyDeduplicationService, contentDeduplicationService)
	deduplicationRuleService := srv.NewDeduplicationRuleService(logger, cmdable, messageTemplateUseCase, deduplicationManager)
	taskService := service.NewTaskService(discardMessageService, shieldService, deduplicationRuleService)
	rabbitmqServer := server.NewMqServer(confData, logger, broker, taskExecutor, handleManager, taskService)
	cronTask := server.NewCronServer(logger, mqHelperMqHelper, cmdable)
	app := newApp(logger, rabbitmqServer, cronTask)
	return app, func() {
		cleanup()
	}, nil
}
