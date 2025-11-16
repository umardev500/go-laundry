package repository

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/tenant"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/internal/domain"
)

type TenantRepository interface {
	Find(ctx *core.Context, f *domain.TenantFilter) ([]*domain.Tenant, int, error)
}

type tenantRepositoryImpl struct {
	client *db.Client
}

func NewTenantRepository(c *db.Client) TenantRepository {
	return &tenantRepositoryImpl{
		client: c,
	}
}

// Find implements TenantRepository.
func (r *tenantRepositoryImpl) Find(ctx *core.Context, f *domain.TenantFilter) ([]*domain.Tenant, int, error) {
	q := r.client.GetConn(ctx).Tenant.Query()
	criteria := f.Criteria

	// Apply search filter
	if criteria.Search != nil {
		q = q.Where(
			tenant.Or(
				tenant.NameContainsFold(*criteria.Search),
			),
		)
	}

	// Count total before pagination
	totalCount, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	q = q.Offset(f.Pagination.Offset).Limit(f.Pagination.Limit)

	// Apply ordering
	orderStr := string(f.Order.Field)
	if f.Order.Dir == core.DESC {
		q = q.Order(ent.Desc(orderStr))
	} else {
		q = q.Order(ent.Asc(orderStr))
	}

	results, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return r.mapEntToDomainList(results), totalCount, nil
}

// --- Helpers ---
func (r *tenantRepositoryImpl) mapEntToDomain(t *ent.Tenant) *domain.Tenant {
	if t == nil {
		return nil
	}

	return &domain.Tenant{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (r *tenantRepositoryImpl) mapEntToDomainList(tenants []*ent.Tenant) []*domain.Tenant {
	var result []*domain.Tenant
	for _, t := range tenants {
		result = append(result, r.mapEntToDomain(t))
	}
	return result
}
