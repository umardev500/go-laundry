package auth

import (
	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(
	NewAuthHandler,
)
