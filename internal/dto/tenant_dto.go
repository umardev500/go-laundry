package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
)

type TenantFilter struct {
	Search *string `query:"search"`

	Limit    int     `query:"limit"`
	Page     int     `query:"page"`
	OrderBy  *string `query:"order_by"`
	OrderDir *string `query:"order_dir"`
}

type Tenant struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// --- Methods ---
func (f *TenantFilter) ToDomain() (*domain.TenantFilter, error) {
	// Ensure a minimuf page of 1
	page := max(f.Page, 1)

	var order *core.Order[domain.TenantOrderField]
	if f.OrderBy != nil {
		dir := core.ASC
		if f.OrderDir != nil {
			dir = core.OrderDirection(*f.OrderDir)
		}

		order = &core.Order[domain.TenantOrderField]{
			Field: domain.TenantOrderField(*f.OrderBy),
			Dir:   dir,
		}

		err := order.Validate(func(tof domain.TenantOrderField) bool {
			return tof == domain.TenantOrderFieldCreatedAt || tof == domain.TenantOrderFieldUpdatedAt
		})
		if err != nil {
			return nil, err
		}
	}

	var paging *core.Pagination
	if f.Limit > 0 {
		paging = &core.Pagination{
			Limit:  f.Limit,
			Offset: (page - 1) * f.Limit,
		}
	}

	filter := domain.NewTenantFilter(
		&domain.TenantFilterCriteria{
			Search: f.Search,
		},
		paging,
		order,
	)

	return &filter, nil
}
