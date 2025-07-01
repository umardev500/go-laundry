package service

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain"
)

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) domain.AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Login implements domain.AuthService.
func (a *authService) Login(ctx context.Context, payload *domain.LoginRequest) (*domain.LoginResponse, error) {
	u, err := a.userRepo.GetByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: "token",
		User:  u,
	}, nil
}
