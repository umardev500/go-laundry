package seed

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/types"
)

// Run executes all registered seeders in order.
func Run(ctx context.Context, seeders []types.Seeder) error {
	log.Info().Msg("🌱 Starting database seeding...")

	for _, s := range seeders {
		if err := s.Seed(ctx); err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Database seeding complete")
	return nil
}
