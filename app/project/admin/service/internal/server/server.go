package server

import (
	"github.com/google/wire"
)

// ProviderSet is serviceName providers.
var ProviderSet = wire.NewSet(NewHTTPServer)
