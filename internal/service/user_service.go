package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Find(ctx *core.Context, f domain.UserFilter) ([]*domain.User, int, error) {
	return s.repo.Find(ctx, f)
}

func (s *UserService) FindByEmail(ctx *core.Context, email string) (*domain.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *UserService) FindByID(ctx *core.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}
