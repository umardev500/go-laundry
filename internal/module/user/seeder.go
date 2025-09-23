package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent/user"
	"github.com/umardev500/go-laundry/internal/db"
	"golang.org/x/crypto/bcrypt"
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
	users := []struct {
		ID       uuid.UUID
		Email    string
		Password string
	}{
		{uuid.MustParse("11111111-1111-1111-1111-111111111111"), "alice@example.com", "password123"},
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), "bob@example.com", "password123"},
		{uuid.MustParse("33333333-3333-3333-3333-333333333333"), "charlie@example.com", "password123"},
	}

	var client = s.client.Client

	for _, u := range users {
		// Hash password
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password for %s: %w", u.Email, err)
		}

		// Skip if user already exists
		exists, err := client.User.Query().Where(user.EmailEQ(u.Email)).Exist(ctx)
		if err != nil {
			return fmt.Errorf("failed to check existing user: %w", err)
		}
		if exists {
			continue
		}

		// Create user with hardcoded ID
		_, err = client.User.Create().
			SetID(u.ID).
			SetEmail(u.Email).
			SetPassword(string(hashed)).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create user %s: %w", u.Email, err)
		}
	}

	return nil
}
