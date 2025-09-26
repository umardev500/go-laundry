package auth

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewSeeder,
	NewHandler,
	NewServiceImpl,
	NewRepository,
	wire.Bind(new(Service), new(*serviceImpl)),
)
