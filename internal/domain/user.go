package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserOrderField string

const (
	CreatedAt UserOrderField = "created_at"
	UpdatedAt UserOrderField = "updated_at"
)

type UserFilter struct {
	Pagination core.Pagination
	Order      core.Order[UserOrderField]
	Search     *string
}

func NewUserFilter(
	search *string,
	pagination *core.Pagination,
	order *core.Order[UserOrderField],
) UserFilter {
	f := UserFilter{}

	f.Search = search

	// Set pagination with default fallback
	if pagination != nil {
		f.Pagination = *pagination
	} else {
		f.Pagination = core.DefaultPagination()
	}

	// Set order with default fallback
	if order != nil {
		f.Order = *order
	} else {
		f.Order = core.DefaultOrder(UpdatedAt)
	}

	return f
}
