package biz

import "github.com/google/wire"

// ProviderSet is initDB providers.
var ProviderSet = wire.NewSet(NewAuthorizationUsecase)
