package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/db"
	"github.com/umardev500/laundry/pkg/security"
)

func SeedUsers(client *db.Client) error {
	log.Info().Msg("User seeding...")
	ctx := context.Background()

	// Hardcoded UUIDs for reference
	aliceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	bobID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	users := []struct {
		ID       uuid.UUID
		Email    string
		Password string
		Profile  struct {
			Name string
		}
	}{
		{
			ID:       aliceID,
			Email:    "alice@example.com",
			Password: "secret123",
			Profile:  struct{ Name string }{Name: "Alice"},
		},
		{
			ID:       bobID,
			Email:    "bob@example.com",
			Password: "password",
			Profile:  struct{ Name string }{Name: "Bob"},
		},
	}

	for _, u := range users {
		// Hash password
		hashedPassword, err := security.HashPassword(u.Password)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to hash password for user %s", u.Email)
			return err
		}

		// Create or update user
		err = client.GetConn(ctx).User.Create().
			SetID(u.ID).
			SetEmail(u.Email).
			SetPassword(string(hashedPassword)).
			OnConflict(
				sql.ConflictColumns("id"),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to seed user %s", u.Email)
			return err
		}

		// 2. Create or upsert profile separately
		err = client.GetConn(ctx).Profile.Create().
			SetUserID(u.ID).
			SetName(u.Profile.Name).
			OnConflict(sql.ConflictColumns("user_id")).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to seed profile for user %s", u.Email)
			return err
		}
	}

	return nil
}
