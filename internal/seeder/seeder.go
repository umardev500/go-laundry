package seeder

import (
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/db"
)

func RunAll(client *db.Client) error {
	log.Info().Msg("Running seeder")
	if err := SeedUsers(client); err != nil {
		return err
	}
	return nil
}
