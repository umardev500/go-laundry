package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type Service struct {
	repo user.Repository
}

func NewService(repo user.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) UpdateUserProfile(ctx context.Context, userID uuid.UUID, u *user.UserProfileUpdate) (*user.UserProfile, error) {
	return s.repo.UpdateUserProfile(ctx, userID, u)
}
