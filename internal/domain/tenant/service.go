package tenant

import "context"

type Service interface {
	CreateTenant(ctx context.Context, tenant *TenantCreate) (*Tenant, error)
}
