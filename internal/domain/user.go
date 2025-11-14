package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
)

// Profile represents the profile information of a user.
// It is nested under User and contains user-specific metadata.
type Profile struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User represents a system user with profile information.
type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Profile   *Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserOrderFields defines the allowed fields to order users by.
type UserOrderField string

const (
	CreatedAt UserOrderField = "created_at"
	UpdatedAt UserOrderField = "updated_at"
)

// UserFilterCriteria defines the filtering options for querying users.
type UserFilterCriteria struct {
	Search         *string
	IncludeProfile bool
}

// UserFilter is used for querying users with pagination, sorting, and filtering.
type UserFilter struct {
	Pagination core.Pagination
	Order      core.Order[UserOrderField]
	Filter     *UserFilterCriteria
}

func NewUserFilter(
	filter *UserFilterCriteria,
	pagination *core.Pagination,
	order *core.Order[UserOrderField],
) UserFilter {
	f := UserFilter{
		Filter: filter,
	}

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
