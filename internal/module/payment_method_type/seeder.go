package paymentmethodtype

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
)

type Seeder struct {
	client *db.Client
}

func NewSeeder(client *db.Client) *Seeder {
	return &Seeder{
		client: client,
	}
}

func (s *Seeder) Seed(ctx context.Context) error {
	defaults := []struct {
		ID          uuid.UUID
		Name        string
		DisplayName string
	}{
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("11111111-1111-1111-1111-111111111111")
				return id
			}(),
			Name:        "credit_card",
			DisplayName: "Credit Card",
		},
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("22222222-2222-2222-2222-222222222222")
				return id
			}(),
			Name:        "bank_transfer",
			DisplayName: "Bank Transfer",
		},
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("33333333-3333-3333-3333-333333333333")
				return id
			}(),
			Name:        "cash",
			DisplayName: "Cash",
		},
	}

	fmt.Println("Seeding payment method types...")

	client := s.client.Client

	for _, d := range defaults {
		err := client.PaymentMethodType.
			Create().
			SetID(d.ID).
			SetName(d.Name).
			SetDisplayName(d.DisplayName).
			OnConflict(
				sql.ConflictColumns("name"),
			).
			UpdateNewValues().
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}
