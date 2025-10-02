package category

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `json:"id"`
	TenantID    *uuid.UUID `json:"tenant_id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Create struct {
	TenantID    *uuid.UUID
	Name        string
	Description *string
}

type Update struct {
	Name        *string
	Description *string
}

type Filter struct {
	Query   string  `query:"query"`
	Offset  int     `query:"offset"`
	Limit   int     `query:"limit"`
	OrderBy OrderBy `query:"order_by"`
}

func (f Filter) WithDefaults() Filter {
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
	OrderByNameAsc  OrderBy = "name_asc"
	OrderByNameDesc OrderBy = "name_desc"
)
