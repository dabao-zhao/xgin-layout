package handlers

import "github.com/google/wire"

// ProviderSet is handler providers.
var ProviderSet = wire.NewSet(NewGreeterHandler)
