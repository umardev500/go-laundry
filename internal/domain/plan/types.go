package plan

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	MaxOrders   *int      `json:"max_orders"`
	MaxUsers    *int      `json:"max_users"`
	Price       *float64  `json:"price"`
	Duration    *int      `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PlanFilter struct {
	IncludeDeleted bool
}

func (f PlanFilter) WithDefaults() PlanFilter {
	return f
}
