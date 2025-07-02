package usecase

import (
	"context"
	"errors"

	"github.com/umardev500/go-laundry/internal/domain"
	sharedjwt "github.com/umardev500/go-laundry/pkg/jwt"
	"github.com/umardev500/go-laundry/pkg/transaction"
)

type MerchantRegistrationUsecase struct {
	tm           *transaction.TransactionManager
	userRepo     domain.UserRepository
	merchantRepo domain.MerchantRepository
}

func NewMerchantRegisterUsecase(
	tm *transaction.TransactionManager,
	userRepo domain.UserRepository,
	merchantRepo domain.MerchantRepository,
) domain.MerchantUsecase {
	return &MerchantRegistrationUsecase{
		tm:           tm,
		userRepo:     userRepo,
		merchantRepo: merchantRepo,
	}
}

func (u *MerchantRegistrationUsecase) Register(ctx context.Context, merchant *domain.CreateMerchantRequest) error {
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

	// Using transaction
	err = u.tm.WithTx(ctx, func(ctx context.Context) error {
		// Create merchant
		merchantInput := &domain.CreateMerchantInput{
			Name:    merchant.Name,
			Email:   merchant.Email,
			Phone:   merchant.Phone,
			Address: merchant.Address,
		}
		merchantCreated, err := u.merchantRepo.Create(ctx, merchantInput)
		if err != nil {
			return err
		}

		// Set merchant id to the user
		if err := u.userRepo.SetMerchantID(ctx, claims.Sub, merchantCreated.ID); err != nil {
			return err
		}

		// TODO: Assign 'Owner' role to user

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
