package paymentmethod

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"
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
		ID                  uuid.UUID
		PaymentMethodTypeID uuid.UUID
		Metadata            any
	}{
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("11111111-1111-1111-1111-111111111111")
				return id
			}(),
			PaymentMethodTypeID: func() uuid.UUID {
				id, _ := uuid.Parse("11111111-1111-1111-1111-111111111111")
				return id
			}(),
			Metadata: paymentmethodtype.QRISMetadata{
				ImageURL:     "https://example.com/image.png",
				MerchantName: "Example Merchant",
			},
		},
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("22222222-2222-2222-2222-222222222222")
				return id
			}(),
			PaymentMethodTypeID: func() uuid.UUID {
				id, _ := uuid.Parse("22222222-2222-2222-2222-222222222222")
				return id
			}(),
			Metadata: paymentmethodtype.BankStransferMetadata{
				Name:          "John Doe",
				AccountNumber: "1234567890",
				AccountHolder: "John Doe",
			},
		},
		{
			ID: func() uuid.UUID {
				id, _ := uuid.Parse("33333333-3333-3333-3333-333333333333")
				return id
			}(),
			PaymentMethodTypeID: func() uuid.UUID {
				id, _ := uuid.Parse("33333333-3333-3333-3333-333333333333")
				return id
			}(),
			Metadata: paymentmethodtype.CashMetadata{
				Note: func() *string {
					s := "Please pay at the counter"
					return &s
				}(),
			},
		},
	}

	fmt.Println("seeding payment methods...")

	client := s.client.Client

	for _, d := range defaults {
		metaMap, err := structToMap(d.Metadata)
		if err != nil {
			return err
		}

		_, err = client.PaymentMethod.Create().
			SetID(d.ID).
			SetPaymentMethodTypeID(d.PaymentMethodTypeID).
			SetMetadata(metaMap).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func structToMap(v any) (map[string]any, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}
	return m, nil
}
