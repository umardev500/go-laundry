package seeds

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/ent"
)

func SeedUsers(ctx context.Context, client *ent.Tx) error {
	users := []struct {
		ID           uuid.UUID
		Name         string
		Email        string
		PasswordHash string
		MerchantID   uuid.UUID
	}{
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("825728b1-37c9-44fd-b20d-f11e206a45d7")
				return id
			}(),
			Name:         "John Doe",
			Email:        "john@gmail.com",
			PasswordHash: "$2a$10$PMA5ZP8QkqAKC7qBix1b0.7lWGYepsdTa0ltyQTEOa1skfSiR7cbK",
			MerchantID: func() uuid.UUID {
				id, _ := uuid.Parse("60836051-d99f-4ae0-912c-70c0c670358d")
				return id
			}(),
		},
	}

	for _, u := range users {
		_, err := client.User.
			Create().
			SetID(u.ID).
			SetName(u.Name).
			SetEmail(u.Email).
			SetPasswordHash(u.PasswordHash).
			SetMerchantsID(u.MerchantID).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
