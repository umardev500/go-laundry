package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/errors"
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

// NewUser creates a new User domain entity
// It requires a non-nil Profile and enforces all invariants.
func NewUser(email, password string, profile *Profile) (*User, error) {
	if profile == nil {
		return nil, errors.NewErrProfileRequired()
	}

	return &User{
		Email:    email,
		Password: password,
		Profile:  profile,
	}, nil
}

// --- Factories ---

// NewProfile creates a Profile domain entity.
func NewProfile(name string) *Profile {
	return &Profile{
		Name: name,
	}
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

// --- Methods ---
func (u *User) Update(email, password *string) (*User, error) {
	if email != nil {
		u.Email = *email
	}
	if password != nil {
		u.Password = *password
	}

	return u, nil
}

func (p *Profile) Update(name *string) (*Profile, error) {
	if name != nil {
		p.Name = *name
	}

	return p, nil
}
