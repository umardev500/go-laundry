package usecase

import (
	"context"
	"errors"

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

func (u *MerchantRegistrationUsecase) Register(ctx context.Context, merchant *domain.CreateMerchantInput) error {
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

	// Check if user already owns a merchant
	exists, err := u.merchantRepo.ExistsByUserID(ctx, claims.Sub)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already owns a merchant")
	}

	// Create merchant
	merchantInput := &domain.CreateMerchantInput{}
	if _, err := u.merchantRepo.Create(ctx, merchantInput); err != nil {
		return err
	}

	// TODO: Assign 'Owner' role to user

	panic("unimplemented")
}
