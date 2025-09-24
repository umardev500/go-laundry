package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type serviceImpl struct {
	repo user.Repository
}

func NewService(repo user.Repository) user.Service {
	return &serviceImpl{
		repo: repo,
	}
}

// CreateProfile implements user.Service.
func (s *serviceImpl) CreateProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileCreate) (*user.Profile, error) {
	return s.repo.CreateUserProfile(ctx, userID, u)
}

func (s *serviceImpl) CreateUser(ctx context.Context, u *user.UserCreate) (*user.User, error) {
	// Hash password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashBytes)

	return s.repo.CreateUser(ctx, u)
}

func (s *serviceImpl) UpdateUserProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileUpdate) (*user.Profile, error) {
	return s.repo.UpdateUserProfile(ctx, userID, u)
}
