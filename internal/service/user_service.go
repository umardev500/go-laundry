package service

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	sharedjwt "github.com/umardev500/go-laundry/pkg/jwt"
)

type userService struct {
	repo domain.UserRepository
}

func (u *userService) GetAll(ctx context.Context, params *domain.GetUsersParams) ([]*ent.User, int, error) {
	claims, err := sharedjwt.Claims[*domain.Claims](ctx)
	if err != nil {
		return nil, 0, err
	}

	params.MerchantID = claims.Merchant

	return u.repo.GetAll(ctx, params)
}

func (u *userService) Create(ctx context.Context, payload *domain.CreateUserRequest) (*ent.User, error) {
	panic("unimplemented")
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{
		repo: repo,
	}
}
