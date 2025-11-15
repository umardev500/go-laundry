package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/commands"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/errors"
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

func (s *UserService) Create(ctx *core.Context, cmd *commands.CreateUserCmd) (*domain.User, error) {
	// Ensure user is not already registered
	if _, err := s.repo.FindByEmail(ctx, cmd.Email); err == nil {
		return nil, errors.NewUserAlreadyExists(cmd.Email)
	}

	profile := domain.NewProfile(cmd.Profile.Name)
	u, err := domain.NewUser(cmd.Email, cmd.Password, profile)
	if err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, u)
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

func (s *UserService) Update(ctx *core.Context, id uuid.UUID, cmd *commands.UpdateUserCmd) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	updatedUserPayload, err := user.Update(cmd.Email, cmd.Password)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.Update(ctx, updatedUserPayload)
	if err != nil {
		if ent.IsConstraintError(err) && cmd.Email != nil {
			return nil, errors.NewUserAlreadyExists(*cmd.Email)
		}

		return nil, err
	}

	return user, nil
}

// UpdateProfile update a user profile.
// It returns a user object with the profile edges.
func (s *UserService) UpdateProfile(ctx *core.Context, userID uuid.UUID, cmd *commands.UpdateProfileCmd) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	updatedProfilePayload, err := user.Profile.Update(cmd.Name)
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateProfile(ctx, userID, updatedProfilePayload); err != nil {
		return nil, err
	}

	return user, nil
}
