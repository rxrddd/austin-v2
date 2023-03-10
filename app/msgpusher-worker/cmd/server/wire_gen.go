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
	"austin-v2/app/msgpusher-worker/internal/sender/smsScript"
	"austin-v2/app/msgpusher-worker/internal/server"
	"austin-v2/app/msgpusher-worker/internal/service"
	"austin-v2/app/msgpusher-worker/internal/service/deduplication"
	"austin-v2/app/msgpusher-worker/internal/service/limiter"
	"austin-v2/app/msgpusher-worker/internal/service/srv"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	taskExecutor := sender.NewTaskExecutor()
	cmdable := data.NewRedisCmd(confData, logger)
	client := data.NewAsynqClient(confData)
	db := data.NewMysqlCmd(confData, logger)
	dataData, cleanup, err := data.NewData(confData, logger, cmdable, db)
	if err != nil {
		return nil, nil, err
	}
	iSendAccountRepo := data.NewSendAccountRepo(dataData, logger)
	sendAccountUseCase := biz.NewSendAccountUseCase(iSendAccountRepo, logger)
	yunPian := smsScript.NewYunPin(logger, client, sendAccountUseCase)
	aliyunSms := smsScript.NewAliyunSms(logger, client, sendAccountUseCase)
	smsManager := smsScript.NewSmsManager(yunPian, aliyunSms)
	smsHandler := handler.NewSmsHandler(logger, cmdable, smsManager)
	iMsgRecordRepo := data.NewMysqlMsgRecordRepo(dataData, logger)
	emailHandler := handler.NewEmailHandler(logger, cmdable, sendAccountUseCase, iMsgRecordRepo)
	officialAccountHandler := handler.NewOfficialAccountHandler(logger, cmdable, sendAccountUseCase, iMsgRecordRepo)
	dingDingRobotHandler := handler.NewDingDingRobotHandler(logger, cmdable, sendAccountUseCase, iMsgRecordRepo)
	dingDingWorkNoticeHandler := handler.NewDingDingWorkNoticeHandler(logger, cmdable, sendAccountUseCase)
	miniProgramHandler := handler.NewMiniProgramHandler(logger, cmdable, sendAccountUseCase, iMsgRecordRepo)
	handleManager := handler.NewHandleManager(smsHandler, emailHandler, officialAccountHandler, dingDingRobotHandler, dingDingWorkNoticeHandler, miniProgramHandler)
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
	iSmsRecordRepo := data.NewSmsRecordRepo(dataData, logger)
	smsRecordUseCase := biz.NewSmsRecordUseCase(iSmsRecordRepo, logger)
	consumeLogic := server.NewLogic(logger, taskExecutor, handleManager, taskService, smsRecordUseCase, client, cmdable)
	asynqServer := data.NewAsynqServer(confData)
	workerServer := server.NewWorkerServer(confData, consumeLogic, asynqServer)
	scheduler := data.NewAsynqScheduler(confData)
	cronTask := server.NewCronServer(logger, scheduler)
	app := newApp(logger, workerServer, cronTask)
	return app, func() {
		cleanup()
	}, nil
}
