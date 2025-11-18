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
	conn := client.GetConn(ctx)

	fastLaundry := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	ezLaundry := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	// Users
	aliceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	bobID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	tenants := []struct {
		ID    uuid.UUID
		Name  string
		Users []uuid.UUID
	}{
		{
			ID:    fastLaundry,
			Name:  "Fast Laundry",
			Users: []uuid.UUID{aliceID},
		},
		{
			ID:    ezLaundry,
			Name:  "Ez Laundry",
			Users: []uuid.UUID{bobID},
		},
	}

	tx, err := conn.Tx(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to start tenant seed transaction")
		return err
	}

	for _, t := range tenants {

		err := tx.Tenant.
			Create().
			SetID(t.ID).
			SetName(t.Name).
			OnConflict(
				sql.ConflictColumns("id"),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			log.Error().Err(err).Str("tenant", t.Name).Msg("failed to create tenant")
			tx.Rollback()
			return err
		}

		for _, userID := range t.Users {
			_, err := tx.TenantUser.
				Create().
				SetUserID(userID).
				SetTenantID(t.ID).
				Save(ctx)
			if err != nil {
				log.Error().Err(err).
					Str("tenant", t.Name).
					Str("user_id", userID.String()).
					Msg("failed to link user to tenant")

				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("failed to commit tenant seed transaction")
		return err
	}

	log.Info().Msg("Tenant seeding completed successfully.")
	return nil
}
