package types

import (
	"context"
)

// Seeder is a common interface for all module seeders.
type Seeder interface {
	Seed(ctx context.Context) error
}
