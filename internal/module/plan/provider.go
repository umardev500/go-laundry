package plan

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewSeeder,
	NewRepository,
	NewService,
	NewHandler,
)
