package platformuser

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive    Status = "active"
	StatusSuspended Status = "suspended"
	StatusDeleted   Status = "deleted"
)

// PlatformUser entity
type PlatformUser struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	Status    Status     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Create payload
type Create struct {
	UserID uuid.UUID
	Status *string
}

// Update payload
type Update struct {
	Status *string
}

// OrderBy
type OrderBy string

const (
	OrderByCreatedAtAsc  OrderBy = "created_at_asc"
	OrderByCreatedAtDesc OrderBy = "created_at_desc"
	OrderByUpdatedAtAsc  OrderBy = "updated_at_asc"
	OrderByUpdatedAtDesc OrderBy = "updated_at_desc"
)

// Filter for listing
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
		f.OrderBy = "created_at_asc"
	}
	return f
}
