package feature

import (
	"context"
	"fmt"

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
	var client = s.client.Client

	fmt.Println("seeding features...")

	features := []struct {
		Name        string
		Description string
		Permissions []string
	}{
		{
			Name:        "Orders",
			Description: "Manage laundry orders",
			Permissions: []string{
				"create_order",
				"read_order",
				"update_order",
				"delete_order",
			},
		},
	}

	for _, f := range features {
		feat, err := client.Feature.
			Create().
			SetName(f.Name).
			SetDescription(f.Description).
			Save(ctx)
		if err != nil {
			return err
		}

		for _, permName := range f.Permissions {
			_, err := client.Permission.
				Create().
				SetName(permName).
				SetFeature(feat).
				Save(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
