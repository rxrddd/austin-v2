package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

type CronTask struct {
	logger    *log.Helper
	scheduler *asynq.Scheduler
}

func NewCronServer(
	logger log.Logger,
	scheduler *asynq.Scheduler,
) *CronTask {
	return &CronTask{
		logger:    log.NewHelper(log.With(logger, "module", "server/cron")),
		scheduler: scheduler,
	}
}

func (l *CronTask) Start(context.Context) error {
	//每早8点发送被屏蔽的消息
	//l.scheduler.Register("0 0 8 * * ?", newNightShieldTask())

	return l.scheduler.Start()
}
func (l *CronTask) Stop(context.Context) error {
	l.scheduler.Shutdown()
	return nil
}
