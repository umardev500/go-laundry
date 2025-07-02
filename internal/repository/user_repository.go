package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/merchant"
	"github.com/umardev500/go-laundry/internal/ent/user"
	"github.com/umardev500/go-laundry/pkg/transaction"
)

type userRepository struct {
	client *ent.Client
	tm     *transaction.TransactionManager
}

func NewUserRepository(client *ent.Client, tm *transaction.TransactionManager) domain.UserRepository {
	return &userRepository{
		client: client,
		tm:     tm,
	}
}

func (r *userRepository) Create(ctx context.Context, payload *domain.CreateUserInput) (*ent.User, error) {
	return r.client.User.
		Create().
		SetName(payload.Name).
		SetEmail(payload.Email).
		SetPasswordHash(payload.Password).
		SetMerchantID(payload.MerchantID).
		Save(ctx)
}

func (r *userRepository) GetAll(ctx context.Context, params *domain.GetUsersParams) ([]*ent.User, int, error) {
	query := r.client.User.
		Query().
		Where(
			user.HasMerchantWith(
				merchant.IDEQ(params.MerchantID),
			),
		)

	// GEt total count first
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	users, err := query.
		Limit(params.Limit).
		Offset(params.Offset).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.client.User.
		Query().
		WithMerchant().
		Where(user.EmailEQ(email)).
		Only(ctx)
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.client.User.
		Query().
		WithMerchant().
		Where(user.ID(id)).
		Only(ctx)
}

func (r *userRepository) SetMerchantID(ctx context.Context, userID uuid.UUID, merchantID uuid.UUID) error {
	client := r.tm.GetClient(ctx)

	_, err := client.User.
		Update().
		Where(user.ID(userID)).
		SetMerchantID(merchantID).
		Save(ctx)
	return err
}
