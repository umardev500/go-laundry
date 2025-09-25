package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

func (r CreateUserRequest) ToUserCreate(tenantID *uuid.UUID) *user.UserCreate {
	return &user.UserCreate{
		Email:    r.Email,
		Password: r.Password,
		TenantID: tenantID,
	}
}
