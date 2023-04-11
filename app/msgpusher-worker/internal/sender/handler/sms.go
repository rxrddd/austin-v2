package handler

import (
	"austin-v2/app/msgpusher-worker/internal/sender/smsScript"
	"austin-v2/common/enums/channelType"
	"austin-v2/pkg/types"
	"austin-v2/utils/timeHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type SmsHandler struct {
	BaseHandler

	logger     *log.Helper
	rds        redis.Cmdable
	smsManager *smsScript.SmsManager
}

func NewSmsHandler(
	logger log.Logger,
	rds redis.Cmdable,
	smsManager *smsScript.SmsManager,
) *SmsHandler {
	return &SmsHandler{
		logger:     log.NewHelper(log.With(logger, "module", "sender/sms")),
		rds:        rds,
		smsManager: smsManager,
	}
}

func (h *SmsHandler) Name() string {
	return channelType.TypeCodeEn[channelType.Sms]
}

func (h *SmsHandler) Execute(_ context.Context, _ *types.TaskInfo) (err error) {
	fmt.Println("sms sender " + timeHelper.CurrentTimeYMDHIS())
	//多短信渠道发送方式 没有测试号目前没法测试
	//route, _ := h.smsManager.Get(taskInfo.SmsChannel)
	//route, _ := h.smsManager.Get("yunpian")
	//route.Send(ctx,taskInfo)
	return nil
}
