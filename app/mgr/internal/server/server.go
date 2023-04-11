package server

import (
	"github.com/google/wire"
)

// ServerProviderSet is server providers.
var ServerProviderSet = wire.NewSet(
	NewHTTPServer,
	NewWorkerServer,
	//NewCronServer,
)
