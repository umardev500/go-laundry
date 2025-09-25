package seed

import (
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/feature"
	"github.com/umardev500/go-laundry/internal/module/plan"
	"github.com/umardev500/go-laundry/internal/module/user"
	"github.com/umardev500/go-laundry/internal/types"
)

func ProvideSeeders(
	userSeeder *user.Seeder,
	authSeeder *auth.Seeder,
	featureSeeder *feature.Seeder,
	planSeeder *plan.Seeder,
) []types.Seeder {
	return []types.Seeder{
		userSeeder,
		authSeeder,
		featureSeeder,
		planSeeder,
	}
}
