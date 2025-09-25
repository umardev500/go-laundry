package auth

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
	fmt.Println("seeding auth...")
	return nil
}
