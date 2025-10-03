package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/category"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
)

type Services struct {
	ID          uuid.UUID          `json:"id"`
	TenantID    *uuid.UUID         `json:"tenant_id"`
	Tenant      *tenant.Tenant     `json:"tenant"`
	CategoryID  *uuid.UUID         `json:"category_id"`
	Category    *category.Category `json:"category"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	BasePrice   float64            `json:"base_price"`
	UnitID      *uuid.UUID         `json:"unit_id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type Create struct {
	TenantID    *uuid.UUID
	CategoryID  *uuid.UUID
	Name        string
	Description *string
	BasePrice   float64
	UnitID      *uuid.UUID
}

// Update does not allow changing CategoryID (immutable after creation)
type Update struct {
	Name        *string
	Description *string
	BasePrice   *float64
	UnitID      *uuid.UUID
}

type Filter struct {
	Query           string  `query:"query"`
	Offset          int     `query:"offset"`
	Limit           int     `query:"limit"`
	OrderBy         OrderBy `query:"order_by"`
	IncludeTenant   bool    `query:"include_tenant"`
	IncludeCategory bool    `query:"include_category"`
}

func (f *Filter) WithDefaults() *Filter {
	if f.Limit == 0 {
		f.Limit = 10
	}
	if f.Offset == 0 {
		f.Offset = 0
	}
	if f.OrderBy == "" {
		f.OrderBy = OrderByNameAsc
	}
	return f
}

type OrderBy string

const (
	OrderByNameAsc       OrderBy = "name_asc"
	OrderByNameDesc      OrderBy = "name_desc"
	OrderByPriceAsc      OrderBy = "price_asc"
	OrderByPriceDesc     OrderBy = "price_desc"
	OrderByCreatedAtAsc  OrderBy = "created_at_asc"
	OrderByCreatedAtDesc OrderBy = "created_at_desc"
)
