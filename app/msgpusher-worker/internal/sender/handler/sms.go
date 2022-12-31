package handler

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/pkg/types"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type SmsHandler struct {
	BaseHandler

	logger *log.Helper
	rds    redis.Cmdable
}

func NewSmsHandler(
	logger log.Logger,
	rds redis.Cmdable,
) *SmsHandler {
	return &SmsHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/sms")),
		rds:    rds,
	}
}

func (h *SmsHandler) Name() string {
	return channelType.TypeCodeEn[channelType.Sms]
}

func (h *SmsHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	fmt.Println("sms sender")
	return nil
}
