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
	ctx := context.Background()

	// Hardcoded UUIDs for reference
	aliceID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	bobID := uuid.MustParse("22222222-2222-2222-2222-222222222222")

	users := []struct {
		ID       uuid.UUID
		Email    string
		Password string
	}{
		{ID: aliceID, Email: "alice@example.com", Password: "secret123"},
		{ID: bobID, Email: "bob@example.com", Password: "password"},
	}

	for _, u := range users {
		// Hash password using bcrypt
		hashedPassword, err := security.HashPassword(u.Password)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to hash password for user %s", u.Email)
			return err
		}

		// Insert into DB
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
	}

	return nil
}
