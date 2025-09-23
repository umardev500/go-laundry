package user

import (
	"github.com/google/wire"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

var ProviderSet = wire.NewSet(
	NewSeeder,
	NewRepositoryImpl,
	wire.Bind(new(user.Repository), new(*repositoryImpl)),
)
