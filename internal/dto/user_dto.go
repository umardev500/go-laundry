package dto

import (
	"github.com/umardev500/laundry/internal/commands"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
)

// Profile represents the profile metadata associated with a user.
// This structure is typically returned as part of the User response object.
type Profile struct {
	Name string `json:"name"`
}

// User represents a simplified view of a user entity that is exposed via API response.
// It includes basic identifying information and optionally the user's profile.
type User struct {
	Email   string   `json:"email"`
	Profile *Profile `json:"profile"`
}

type UserFilter struct {
	Search         *string `query:"search"`
	IncludeProfile bool    `query:"include_profile"`

	Limit    int     `query:"limit"`
	Page     int     `query:"page"`
	OrderBy  *string `query:"order_by"`
	OrderDir *string `query:"order_dir"`
}

// CreateProfileDTO represenets the expected payload when creating a profile.
// Used in HTTP request bodies to receive profile information from clients.
type CreateProfileDTO struct {
	Name string `json:"name"`
}

// CreateProfileDTO represents the expected payload when creating a new user.
// Typically converted into application commands inside the handler layer.
type CreateUserDTO struct {
	Email    string            `json:"email" validate:"required,email"`
	Password string            `json:"password" validate:"required,min=6"`
	Profile  *CreateProfileDTO `json:"profile"`
}

// UpdateUserDTO represents the expected payload for updateing a user.
type UpdateUserDTO struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"omitempty,password"`
}

// UpdateProfileDTO represents the expexted payload for update a profile of a user.
type UpdateProfileDTO struct {
	Name *string `json:"name" validate:"omitempty,min=3"`
}

// --- Methods ---

func (f *UserFilter) ToDomain() (*domain.UserFilter, error) {
	// Ensure a minimum page of 1
	page := max(f.Page, 1)

	var order *core.Order[domain.UserOrderField]
	if f.OrderBy != nil {
		dir := core.ASC
		if f.OrderDir != nil {
			dir = core.OrderDirection(*f.OrderDir)
		}

		order = &core.Order[domain.UserOrderField]{
			Field: domain.UserOrderField(*f.OrderBy),
			Dir:   dir,
		}

		err := order.Validate(func(uof domain.UserOrderField) bool {
			return uof == domain.CreatedAt || uof == domain.UpdatedAt
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

	filter := domain.NewUserFilter(
		&domain.UserFilterCriteria{
			Search:         f.Search,
			IncludeProfile: f.IncludeProfile,
		},
		paging,
		order,
	)

	return &filter, nil
}

func (c *CreateUserDTO) ToCmd() (*commands.CreateUserCmd, error) {
	return &commands.CreateUserCmd{
		Email:    c.Email,
		Password: c.Password,
		Profile: &commands.CreateProfileCmd{
			Name: c.Profile.Name,
		},
	}, nil
}

func (u *UpdateUserDTO) ToCmd() (*commands.UpdateUserCmd, error) {
	return &commands.UpdateUserCmd{
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (u *UpdateProfileDTO) ToCmd() (*commands.UpdateProfileCmd, error) {
	return &commands.UpdateProfileCmd{
		Name: u.Name,
	}, nil
}
