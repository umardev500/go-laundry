package role

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/permission"
)

type Role struct {
	ID          uuid.UUID                `json:"id"`
	Name        string                   `json:"name"`
	Description *string                  `json:"description"`
	Permissions []*permission.Permission `json:"permissions"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

type RoleCreate struct {
	Name        string
	Description *string
}

type RoleUpdate struct {
	Name        *string
	Description *string
}

type OrderBy string

const (
	OrderByCreatedAtAsc  OrderBy = "created_at_asc"
	OrderByCreatedAtDesc OrderBy = "created_at_desc"
)

type Filter struct {
	Query              string  `query:"query"`
	Limit              int     `query:"limit"`
	Offset             int     `query:"offset"`
	OrderBy            OrderBy `query:"order_by"`
	IncludeDeleted     bool    `query:"include_deleted"`
	IncludePermissions bool    `query:"include_permissions"`
}

func (f Filter) WithDefaults() *Filter {
	if f.Limit == 0 {
		f.Limit = 10 // default page size
	}
	if f.Offset == 0 {
		f.Offset = 0
	}
	if f.OrderBy == "" {
		f.OrderBy = "created_at desc" // default ordering
	}
	return &f
}
