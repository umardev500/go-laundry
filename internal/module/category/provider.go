package category

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewRepositoryImpl,
	NewService,
	NewHandler,
)
