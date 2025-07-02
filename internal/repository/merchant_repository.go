package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/merchant"
	"github.com/umardev500/go-laundry/internal/ent/user"
)

type merchantRepository struct {
	client *ent.Client
}

func (r *merchantRepository) Create(ctx context.Context, payload *domain.CreateMerchantInput) (*ent.Merchant, error) {
	return r.client.Merchant.
		Create().
		SetName(payload.Name).
		SetEmail(payload.Email).
		SetPhone(payload.Phone).
		SetAddress(payload.Address).
		Save(ctx)
}

func (r *merchantRepository) ExistsByUserID(ctx context.Context, ownerID uuid.UUID) (bool, error) {
	exist, err := r.client.Merchant.
		Query().
		Where(
			merchant.HasUsersWith(
				user.IDEQ(ownerID),
			),
		).Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func NewMerchantRepository(client *ent.Client) domain.MerchantRepository {
	return &merchantRepository{
		client: client,
	}
}
