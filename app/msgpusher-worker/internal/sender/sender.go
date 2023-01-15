package sender

import (
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"austin-v2/app/msgpusher-worker/internal/sender/smsScript"
	"github.com/google/wire"
)

// SenderProviderSet is biz providers.
var SenderProviderSet = wire.NewSet(
	handler.NewSmsHandler,
	handler.NewDingDingRobotHandler,
	handler.NewDingDingWorkNoticeHandler,
	handler.NewOfficialAccountHandler,
	handler.NewMiniProgramHandler,
	handler.NewEmailHandler,
	handler.NewHandleManager,
	smsScript.NewYunPin,
	smsScript.NewAliyunSms,
	smsScript.NewSmsManager,
	NewTaskExecutor,
)
