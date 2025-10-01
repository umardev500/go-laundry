package upload

import (
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/domain/upload"
)

var ProvidertSet = wire.NewSet(
	NewHandler,
	NewService,
	wire.Value(upload.BasePath("uploads")),
)
