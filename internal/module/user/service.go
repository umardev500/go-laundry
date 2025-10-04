package user

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
	"golang.org/x/crypto/bcrypt"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type serviceImpl struct {
	repo user.Repository
}

// FindByEmail implements user.Service.
func (s *serviceImpl) FindByEmail(ctx *appContext.ScopedContext, email string) (*user.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

// FindByToken implements user.Service.
func (s *serviceImpl) FindByToken(ctx *appContext.ScopedContext, token string) (*user.User, error) {
	return s.repo.FindByToken(ctx, token)
}

// Update implements user.Service.
func (s *serviceImpl) Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *user.UserUpdate) (*user.User, error) {
	// Hash password if provided
	if payload.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		hashedStr := string(hashed)

		payload.Password = &hashedStr
	}

	// Delegate to repository, passing tenantID for scoping
	return s.repo.Update(ctx, payload, id)
}

// List implements user.Service.
func (s *serviceImpl) List(ctx *appContext.ScopedContext, f *user.Filter) (*types.PageResult[user.User], error) {
	// Apply defaults
	f = f.WithDefaults()

	// Deletegate to repository
	result, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	paginatedResult := utils.Paginate(result.Data, result.Total, f.Offset, f.Limit)
	return paginatedResult, nil
}

func NewService(repo user.Repository) user.Service {
	return &serviceImpl{
		repo: repo,
	}
}

// CreateProfile implements user.Service.
func (s *serviceImpl) CreateProfile(ctx *appContext.ScopedContext, userID uuid.UUID, u *user.ProfileCreate) (*user.Profile, error) {
	return s.repo.CreateProfile(ctx, userID, u)
}

func (s *serviceImpl) Create(ctx *appContext.ScopedContext, u *user.UserCreate) (*user.User, error) {
	// Hash password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashBytes)

	return s.repo.Create(ctx, u)
}

// Delete implements user.Service.
func (s *serviceImpl) Delete(ctx *appContext.ScopedContext, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// Purge implements user.Service.
func (s *serviceImpl) Purge(ctx *appContext.ScopedContext, id uuid.UUID) error {
	return s.repo.PurgeUser(ctx, id)
}

func (s *serviceImpl) UpdateProfile(ctx *appContext.ScopedContext, userID uuid.UUID, u *user.ProfileUpdate) (*user.Profile, error) {
	return s.repo.UpdateProfile(ctx, userID, u)
}
