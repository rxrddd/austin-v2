package biz

import (
	"austin-v2/utils/timeHelper"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
)

type ConsumeLogic struct {
	loggerHelper *log.Helper
	logger       log.Logger
	cli          *asynq.Client
	rds          redis.Cmdable
}

func NewConsumeLogic(
	logger log.Logger,
	cli *asynq.Client,
	rds redis.Cmdable,
) *ConsumeLogic {
	return &ConsumeLogic{
		loggerHelper: log.NewHelper(log.With(logger, "module", "ConsumeLogic")),
		logger:       logger,
		cli:          cli,
		rds:          rds,
	}
}

func (m *ConsumeLogic) Test(msg []byte) error {
	fmt.Println(timeHelper.CurrentTimeYMDHIS())
	return nil
}
