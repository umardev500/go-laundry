package registration

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewHandler,
	NewService,
)
