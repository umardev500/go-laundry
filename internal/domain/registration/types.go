package registration

import (
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type RegisterInput struct {
	Tenant  *tenant.TenantCreate
	Profile *user.ProfileCreate
	User    *user.UserCreate
}
