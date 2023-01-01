package sender

import (
	"austin-v2/app/msgpusher-worker/internal/sender/handler"
	"github.com/google/wire"
)

// SenderProviderSet is biz providers.
var SenderProviderSet = wire.NewSet(
	handler.NewSmsHandler,
	handler.NewEmailHandler,
	NewHandleManager,
	NewTaskExecutor,
)
