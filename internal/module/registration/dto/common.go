package dto

import "github.com/umardev500/go-laundry/internal/domain/user"

type Creds struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

func (c Creds) ToCreds() *user.UserCreate {
	return &user.UserCreate{
		Email:    c.Email,
		Password: c.Password,
	}
}

type Profile struct {
	Name    string  `json:"name" validate:"required"`
	Avatar  *string `json:"avatar"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}
