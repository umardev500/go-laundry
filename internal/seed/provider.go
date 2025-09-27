package seed

import (
	"github.com/umardev500/go-laundry/internal/module/auth"
	"github.com/umardev500/go-laundry/internal/module/feature"
	paymentmethod "github.com/umardev500/go-laundry/internal/module/payment_method"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/module/payment_method_type"
	"github.com/umardev500/go-laundry/internal/module/plan"
	"github.com/umardev500/go-laundry/internal/module/user"
	"github.com/umardev500/go-laundry/internal/types"
)

func ProvideSeeders(
	userSeeder *user.Seeder,
	authSeeder *auth.Seeder,
	featureSeeder *feature.Seeder,
	planSeeder *plan.Seeder,
	paymentMethodTypeSeeder *paymentmethodtype.Seeder,
	paymentMethodSeder *paymentmethod.Seeder,
) []types.Seeder {
	return []types.Seeder{
		userSeeder,
		authSeeder,
		featureSeeder,
		planSeeder,
		paymentMethodTypeSeeder,
		paymentMethodSeder,
	}
}
