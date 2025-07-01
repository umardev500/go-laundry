package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/user"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) domain.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Create(ctx context.Context, payload *domain.CreateUserInput) (*ent.User, error) {
	r.client.User.
		Create().
		SetMerchantsID(payload.MerchantID).
		Save(ctx)
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
}
