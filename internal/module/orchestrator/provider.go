package orchestrator

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewPaymentService,
)
