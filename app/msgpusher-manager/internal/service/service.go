package service

import "github.com/google/wire"

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(NewMsgPusherManagerService)
