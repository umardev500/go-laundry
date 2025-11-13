package repository

import "github.com/google/wire"

var Set = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(UserRepository), new(*userRepositoryImpl)),
)
