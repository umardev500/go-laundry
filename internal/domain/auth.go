package domain

import (
	"context"

	"github.com/umardev500/go-laundry/internal/ent"
)

// DTO

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *ent.User
}

type AuthService interface {
	Login(ctx context.Context, payload *LoginRequest) (*LoginResponse, error)
}
