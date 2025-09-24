package tenant

import "context"

type Repository interface {
	CreateTenant(ctx context.Context, tenant *TenantCreate) (*Tenant, error)
}
