package seed

import (
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/user"
	"github.com/umardev500/go-laundry/internal/types"
)

func ProvideSeeders(
	userSeeder *user.Seeder,
	authSeeder *auth.Seeder,
) []types.Seeder {
	return []types.Seeder{
		userSeeder,
		authSeeder,
	}
}
