package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
)

type TenantUserFilter struct {
	Search         *string `query:"search"`
	IncludeUser    *bool   `query:"include_user"`
	IncludeProfile *bool   `query:"include_profile"`

	Limit    int     `query:"limit"`
	Page     int     `query:"page"`
	OrderBy  *string `query:"order_by"`
	OrderDir *string `query:"order_dir"`
}

type TenantUser struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// --- Methods ---
func (f *TenantUserFilter) ToDomain() (*domain.TenantUserFilter, error) {
	f.SetDefaults()

	var order *core.Order[domain.TenantUserOrderField]
	if f.OrderBy != nil {
		dir := core.ASC
		if f.OrderDir != nil {
			dir = core.OrderDirection(*f.OrderDir)
		}

		order = &core.Order[domain.TenantUserOrderField]{
			Field: domain.TenantUserOrderField(*f.OrderBy),
			Dir:   dir,
		}

		err := order.Validate(func(tof domain.TenantUserOrderField) bool {
			return tof == domain.TenantUserOrderFieldCreatedAt || tof == domain.TenantUserOrderFieldUpdatedAt
		})
		if err != nil {
			return nil, err
		}
	}

	var paging *core.Pagination
	if f.Limit > 0 {
		paging = &core.Pagination{
			Limit:  f.Limit,
			Offset: (f.Page - 1) * f.Limit,
		}
	}

	filter := domain.NewTenantUserFilter(
		&domain.TenantUserFilterCriteria{
			Search:         f.Search,
			IncludeUser:    *f.IncludeUser,
			IncludeProfile: *f.IncludeProfile,
		},
		paging,
		order,
	)

	return &filter, nil
}

func (f *TenantUserFilter) SetDefaults() {
	if f.IncludeUser == nil {
		t := true
		f.IncludeUser = &t
	}
	if f.IncludeProfile == nil {
		t := true
		f.IncludeProfile = &t
	}

	// Page and Limit
	if f.Page < 1 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 20
	}
}
