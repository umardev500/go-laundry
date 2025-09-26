package plan

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/pkg/utils"
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
	client := s.client.Client

	fmt.Println("seeding plans...")

	plans := []struct {
		Name         string
		MaxOrders    *int
		MaxUsers     *int
		Price        *float64
		DurationDays *int
	}{
		{
			Name:         "Free",
			MaxOrders:    utils.Ptr(5),
			MaxUsers:     utils.Ptr(1),
			Price:        utils.Ptr(0.0),
			DurationDays: utils.Ptr(30),
		},
		{
			Name:         "Basic",
			MaxOrders:    utils.Ptr(10),
			MaxUsers:     utils.Ptr(1),
			Price:        utils.Ptr(9.99),
			DurationDays: utils.Ptr(30),
		},
		{
			Name:         "Standard",
			MaxOrders:    utils.Ptr(50),
			MaxUsers:     utils.Ptr(5),
			Price:        utils.Ptr(19.99),
			DurationDays: utils.Ptr(30),
		},
		{
			Name:         "Premium",
			MaxOrders:    utils.Ptr(0), // unlimited
			MaxUsers:     utils.Ptr(0), // unlimited
			Price:        utils.Ptr(29.99),
			DurationDays: utils.Ptr(30),
		},
	}

	for _, p := range plans {
		// Use Upsert to avoid duplicates if the seeder is run multiple times
		err := client.Plan.
			Create().
			SetID(uuid.New()).
			SetNillableMaxOrders(p.MaxOrders).
			SetNillableMaxUsers(p.MaxUsers).
			SetNillablePrice(p.Price).
			SetNillableDurationDays(p.DurationDays).
			SetName(p.Name).
			OnConflict(
				sql.ConflictColumns("name"),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to seed plan %s: %w", p.Name, err)
		}
		fmt.Printf("Seeded plan: %s\n", p.Name)
	}

	return nil
}
