package handler

import (
	"austin-v2/app/msgpusher-worker/internal/enums/channelType"
	"austin-v2/pkg/types"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type EmailHandler struct {
	BaseHandler

	logger *log.Helper
	rds    redis.Cmdable
}

func NewEmailHandler(
	logger log.Logger,
	rds redis.Cmdable,
) *EmailHandler {
	return &EmailHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/sms")),
		rds:    rds,
	}
}

func (h *EmailHandler) Name() string {
	return channelType.TypeCodeEn[channelType.Email]
}

func (h *EmailHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	fmt.Println("sms sender")
	return nil
}
