package tenant

import (
	"context"

	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
)

type repositoryImpl struct {
	client *db.Client
}

func (r *repositoryImpl) CreateTenant(ctx context.Context, t *tenant.TenantCreate) (*tenant.Tenant, error) {
	conn := r.client.GetConn(ctx)

	tenantReturned, err := conn.Tenant.
		Create().
		SetName(t.Name).
		SetPhone(t.Phone).
		SetEmail(t.Email).
		SetAddress(t.Address).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var result tenant.Tenant
	r.mapFromEnt(tenantReturned, &result)
	return &result, nil
}

func (r *repositoryImpl) mapFromEnt(e *ent.Tenant, to *tenant.Tenant) {
	if to == nil {
		return
	}

	to.ID = e.ID
	to.Name = *e.Name
	to.Phone = *e.Phone
	to.Email = *e.Email
	to.Address = *e.Address
	to.CreatedAt = e.CreatedAt
	to.UpdatedAt = e.UpdatedAt
}

func NewRepository(client *db.Client) tenant.Repository {
	return &repositoryImpl{
		client: client,
	}
}
