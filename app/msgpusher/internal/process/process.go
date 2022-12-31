package process

import (
	"github.com/google/wire"
)

// ProcessProviderSet is server providers.
var ProcessProviderSet = wire.NewSet(
	NewPreParamCheckAction,
	NewAssembleAction,
	NewAfterParamCheckAction,
	NewSendMqAction,
	NewBusinessProcess,
)
