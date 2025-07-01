package service

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain"
	"golang.org/x/crypto/bcrypt"
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

	inputPass := []byte(payload.Password)
	hashedPass := []byte(u.PasswordHash)

	err = bcrypt.CompareHashAndPassword(hashedPass, inputPass)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: "token",
	}, nil
}
