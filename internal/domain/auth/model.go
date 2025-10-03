package auth

import (
	platformuser "github.com/umardev500/go-laundry/internal/domain/platform_user"
	tenantuser "github.com/umardev500/go-laundry/internal/domain/tenant_user"
)

// LoginResolution is a domain object that models the possible
// login contexts for a user.
// - PlatformUser: means the user is a platform-scoped account.
// - TenantUsers: means the user belongs to one or more tenants.
// - Both: means the user must select between platform vs tenant login.
type LoginResolution struct {
	PlatformUser *platformuser.PlatformUser
	TenantUsers  []*tenantuser.TenantUser
}

// IsAmbiguous returns true if user has both a platform account
// and one or more tenant memberships (requires user to choose).
func (r *LoginResolution) IsAmbiguous() bool {
	return r.PlatformUser != nil && len(r.TenantUsers) > 0
}

// IsTenantMulti returns true if user has multiple tenant memberships.
func (r *LoginResolution) IsTenantMulti() bool {
	return len(r.TenantUsers) > 1 && r.PlatformUser == nil
}

// IsSingleTenant returns true if user has exactly one tenant membership.
func (r *LoginResolution) IsSingleTenant() bool {
	return len(r.TenantUsers) == 1 && r.PlatformUser == nil
}

// IsPlatformOnly returns true if only platform user exists.
func (r *LoginResolution) IsPlatformOnly() bool {
	return r.PlatformUser != nil && len(r.TenantUsers) == 0
}
