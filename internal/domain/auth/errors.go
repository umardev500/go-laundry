package auth

import "errors"

var (
	ErrMultipleAccountTypes = errors.New("user belongs to both platform and tenant, must select account type")
	ErrMultipleTenants      = errors.New("user belongs to multiple tenants, must select which tenant to login")
)
