package dto

import (
	"github.com/umardev500/go-laundry/internal/domain/registration"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type CreateUserRequest struct {
	Profile Profile `json:"profile" validate:"required"`
	Creds   Creds   `json:"user" validate:"required"`
}

func (r *CreateUserRequest) ToUserCreate() *registration.CreateUser {
	return &registration.CreateUser{
		Profile: &user.ProfileCreate{
			Name:   r.Profile.Name,
			Avatar: r.Profile.Avatar,
			Phone:  r.Profile.Phone,
		},
		User: r.Creds.ToCreds(),
	}
}
