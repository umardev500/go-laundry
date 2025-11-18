package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
)

type TenantUserOrderField string

const (
	TenantUserOrderFieldCreatedAt TenantUserOrderField = "created_at"
	TenantUserOrderFieldUpdatedAt TenantUserOrderField = "updated_at"
)

type TenantUserFilterCriteria struct {
	Search         *string
	IncludeUser    bool
	IncludeProfile bool
}

type TenantUserFilter struct {
	Pagination core.Pagination
	Order      core.Order[TenantUserOrderField]
	Criteria   *TenantUserFilterCriteria
}

type UserOfTenant struct {
	Name  string
	Email string
}

type TenantUser struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	UserID    uuid.UUID
	User      UserOfTenant
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTenantUser creates a new Tenant User domain entity.
func NewTenantUser(tenantID, userID uuid.UUID) (*TenantUser, error) {
	return &TenantUser{
		TenantID: tenantID,
		UserID:   userID,
	}, nil
}

// NewTenantFilter creates a a new tenant domain filter.
func NewTenantUserFilter(
	criteria *TenantUserFilterCriteria,
	pagination *core.Pagination,
	order *core.Order[TenantUserOrderField],
) TenantUserFilter {
	f := TenantUserFilter{
		Criteria: criteria,
	}

	// Set pagination with default fallback
	f.Pagination = core.DefaultPaginationFallback(pagination)

	// Set order with default fallback
	f.Order = core.DefaultOrderFallback(order, TenantUserOrderFieldCreatedAt)

	return f
}
