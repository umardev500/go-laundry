package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/ent"
)

// DTO

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Type-safe JWT Claims
type Claims struct {
	Sub      uuid.UUID `json:"sub"`
	Merchant uuid.UUID `json:"merchant"`
	jwt.RegisteredClaims
}

// Interface
type AuthService interface {
	Login(ctx context.Context, payload *LoginRequest) (*LoginResponse, error)
	Me(ctx context.Context) (*ent.User, error)
}
