package seeder

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/db"
)

func RunAll(client *db.Client) error {
	log.Info().Msg("Running seeder")

	return client.WithTransaction(context.Background(), func(ctx context.Context) error {
		if err := SeedUsers(client); err != nil {
			return err
		}

		if err := SeedTenants(client); err != nil {
			return err
		}

		return nil
	})
}
