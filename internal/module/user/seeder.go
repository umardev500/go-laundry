package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/ent/platformuser"
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
	client := s.client.Client

	fmt.Println("seeding users...")

	users := []struct {
		ID         uuid.UUID
		Email      string
		Password   string
		Name       string
		IsPlatform bool
	}{
		{uuid.MustParse("11111111-1111-1111-1111-111111111111"), "alice@example.com", "password123", "Alice", true},
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), "bob@example.com", "password123", "Bob", true},
		{uuid.MustParse("33333333-3333-3333-3333-333333333333"), "charlie@example.com", "password123", "Charlie", false}, // end-user only
	}

	for _, u := range users {
		// Skip if user already exists
		userEntity, err := client.User.Query().Where(user.EmailEQ(u.Email)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return fmt.Errorf("failed to check existing user: %w", err)
		}

		if userEntity == nil {
			// Hash password
			hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			if err != nil {
				return fmt.Errorf("failed to hash password for %s: %w", u.Email, err)
			}

			userEntity, err = client.User.Create().
				SetID(u.ID).
				SetEmail(u.Email).
				SetPassword(string(hashed)).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create user %s: %w", u.Email, err)
			}

			_, err = client.Profile.Create().
				SetName(u.Name).
				SetUser(userEntity).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		// If flagged as a platform user → ensure PlatformUser exists
		if u.IsPlatform {
			exists, err := client.PlatformUser.Query().
				Where(platformuser.HasUserWith(user.IDEQ(userEntity.ID))).
				Exist(ctx)
			if err != nil {
				return err
			}
			if !exists {
				_, err = client.PlatformUser.Create().
					SetUser(userEntity).
					SetStatus(platformuser.StatusActive).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed to create platform user for %s: %w", u.Email, err)
				}
			}
		}
	}

	return nil
}
