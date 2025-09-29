package paymentmethodtype

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewSeeder,
	NewHandler,
	NewService,
	NewRepository,
)
