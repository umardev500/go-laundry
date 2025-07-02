package usecase

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain"
	sharedjwt "github.com/umardev500/go-laundry/pkg/jwt"
)

type MerchantRegistrationUsecase struct {
	userRepo     domain.UserRepository
	merchantRepo domain.MerchantRepository
}

func NewMerchantRegisterUsecase(
	userRepo domain.UserRepository,
	merchantRepo domain.MerchantRepository,
) *MerchantRegistrationUsecase {
	return &MerchantRegistrationUsecase{
		userRepo:     userRepo,
		merchantRepo: merchantRepo,
	}
}

func (u *MerchantRegistrationUsecase) Register(ctx context.Context) error {
	// Get claims
	claims, err := sharedjwt.Claims[*domain.Claims](ctx)
	if err != nil {
		return err
	}

	// Get user
	_, err = u.userRepo.GetByID(ctx, claims.Sub)
	if err != nil {
		return err
	}

	// TODO: Check if user already owns a merchant

	// TODO: Create merchant

	// TODO: Assign 'Owner' role to user

	panic("unimplemented")
}
