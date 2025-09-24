package dto

import (
	"github.com/umardev500/go-laundry/internal/domain/registration"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type RegisterRequest struct {
	Tenant  TenantInfo         `json:"tenant" validate:"required,dive"`
	Profile user.ProfileCreate `json:"profile" validate:"required,dive"`
	Creds   Creds              `json:"user" validate:"required,dive"`
}

func (r *RegisterRequest) ToRegisterInput() *registration.RegisterInput {
	return &registration.RegisterInput{
		Tenant:  r.Tenant.ToTenantCreate(),
		Profile: &r.Profile,
		User:    r.Creds.ToCreds(),
	}
}

type TenantInfo struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Address string `json:"address" validate:"required"`
}

func (t TenantInfo) ToTenantCreate() *tenant.TenantCreate {
	return &tenant.TenantCreate{
		Name:    t.Name,
		Phone:   t.Phone,
		Email:   t.Email,
		Address: t.Address,
	}
}

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
