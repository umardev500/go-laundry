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

// List implements user.Service.
func (s *serviceImpl) List(ctx context.Context, filter user.UserFilter) ([]*user.User, error) {
	// Apply defaults
	filter = filter.WithDefaults()

	// Deletegate to repository
	return s.repo.List(ctx, filter)
}

func NewService(repo user.Repository) user.Service {
	return &serviceImpl{
		repo: repo,
	}
}

// CreateProfile implements user.Service.
func (s *serviceImpl) CreateProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileCreate) (*user.Profile, error) {
	return s.repo.CreateProfile(ctx, userID, u)
}

func (s *serviceImpl) Create(ctx context.Context, u *user.UserCreate) (*user.User, error) {
	// Hash password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashBytes)

	return s.repo.Create(ctx, u)
}

// Delete implements user.Service.
func (s *serviceImpl) Delete(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error {
	return s.repo.Delete(ctx, tenantID, userID)
}

// Purge implements user.Service.
func (s *serviceImpl) Purge(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error {
	return s.repo.PurgeUser(ctx, tenantID, userID)
}

func (s *serviceImpl) UpdateProfile(ctx context.Context, userID uuid.UUID, u *user.ProfileUpdate) (*user.Profile, error) {
	return s.repo.UpdateProfile(ctx, userID, u)
}
