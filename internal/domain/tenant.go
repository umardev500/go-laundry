package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
)

type TenantOrderField string

const (
	TenantOrderFieldCreatedAt TenantOrderField = "created_at"
	TenantOrderFieldUpdatedAt TenantOrderField = "updated_at"
)

type TenantFilterCriteria struct {
	Search *string
}

type TenantFilter struct {
	Pagination core.Pagination
	Order      core.Order[TenantOrderField]
	Criteria   *TenantFilterCriteria
}

type Tenant struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTenant creates a new Tenant domain entity.
func NewTenant(name string) (*Tenant, error) {
	return &Tenant{
		Name: name,
	}, nil
}

// NewTenantFilter creates a a new tenant domain filter.
func NewTenantFilter(
	criteria *TenantFilterCriteria,
	pagination *core.Pagination,
	order *core.Order[TenantOrderField],
) TenantFilter {
	f := TenantFilter{
		Criteria: criteria,
	}

	// Set pagination with default fallback
	f.Pagination = core.DefaultPaginationFallback(pagination)

	// Set order with default fallback
	f.Order = core.DefaultOrderFallback(order, TenantOrderFieldCreatedAt)

	return f
}

// Update updates the tenant fields.
func (t *Tenant) Update(name *string) {
	if name != nil {
		t.Name = *name
	}
}
