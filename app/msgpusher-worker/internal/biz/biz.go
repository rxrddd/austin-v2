package biz

import "github.com/google/wire"

// BizProviderSet is initDB providers.
var BizProviderSet = wire.NewSet(
	NewMessageTemplateUseCase,
	NewSendAccountUseCase,
	NewSmsRecordUseCase,
)
