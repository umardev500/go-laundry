package subscription

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewRepository,
	NewService,
	NewHandler,
)
