package sender

import "github.com/google/wire"

// SenderProviderSet is biz providers.
var SenderProviderSet = wire.NewSet(
	NewHandle,
	NewTaskExecutor,
)
