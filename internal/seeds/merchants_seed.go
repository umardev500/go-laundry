package seeds

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/merchant"
)

func SeedMerchants(ctx context.Context, client *ent.Tx) error {
	merchants := []struct {
		ID      uuid.UUID
		Name    string
		Email   string
		Phone   string
		Address string
	}{
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("60836051-d99f-4ae0-912c-70c0c670358d")
				return id
			}(),
			Name:    "Alpha Laundry",
			Email:   "contact@alphalaundry.com",
			Phone:   "1234567890",
			Address: "123 Alpha Street",
		},
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("c21c88df-86c5-4911-a953-93ade82228bd")
				return id
			}(),
			Name:    "Beta Laundry",
			Email:   "info@betalaundry.com",
			Phone:   "0987654321",
			Address: "456 Beta Avenue",
		},
	}

	for _, m := range merchants {
		// Delete any existing merchant with the same email
		if _, err := client.Merchant.
			Delete().
			Where(merchant.EmailEQ(m.Email)).
			Exec(ctx); err != nil {
			return fmt.Errorf("failed to delete existing merchant %q: %w", m.Email, err)
		}

		// Insert the new merchant
		if _, err := client.Merchant.
			Create().
			SetID(m.ID).
			SetName(m.Name).
			SetEmail(m.Email).
			SetPhone(m.Phone).
			SetAddress(m.Address).
			Save(ctx); err != nil {
			return fmt.Errorf("failed to create merchant %q: %w", m.Name, err)
		}
	}

	return nil
}
