package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/db"
)

func SeedTenants(client *db.Client) error {
	log.Info().Msg("Tenant seeding...")
	ctx := context.Background()
	fastLaundry := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	ezLaundry := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	tenants := []struct {
		ID   uuid.UUID
		Name string
	}{
		{
			ID:   fastLaundry,
			Name: "Fast Laundry",
		},
		{
			ID:   ezLaundry,
			Name: "Ez Laundry",
		},
	}

	for _, t := range tenants {
		err := client.GetConn(ctx).Tenant.Create().
			SetID(t.ID).
			SetName(t.Name).
			OnConflict(
				sql.ConflictColumns("id"),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to seed tenant %s", t.Name)
			return err
		}
	}

	return nil
}
