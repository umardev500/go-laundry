package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/ent/platformuser"
	"github.com/umardev500/go-laundry/ent/tenantuser"
	"github.com/umardev500/go-laundry/ent/user"
	"github.com/umardev500/go-laundry/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type Seeder struct {
	client *db.Client
}

func NewSeeder(client *db.Client) *Seeder {
	return &Seeder{client: client}
}

func (s *Seeder) Seed(ctx context.Context) error {
	client := s.client.Client
	fmt.Println("seeding users...")

	// ✅ Seed tenants first (with required fields)
	tenant1, err := client.Tenant.Create().
		SetID(uuid.MustParse("aaaa1111-1111-1111-1111-111111111111")).
		SetName("Tenant One").
		SetPhone("123-456-7890").
		SetEmail("tenant1@example.com").
		SetAddress("123 Main St").
		Save(ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed to create tenant1: %w", err)
	}
	if ent.IsConstraintError(err) {
		tenant1, _ = client.Tenant.Get(ctx, uuid.MustParse("aaaa1111-1111-1111-1111-111111111111"))
	}

	tenant2, err := client.Tenant.Create().
		SetID(uuid.MustParse("aaaa2222-2222-2222-2222-222222222222")).
		SetName("Tenant Two").
		SetPhone("987-654-3210").
		SetEmail("tenant2@example.com").
		SetAddress("456 Elm St").
		Save(ctx)
	if err != nil && !ent.IsConstraintError(err) {
		return fmt.Errorf("failed to create tenant2: %w", err)
	}
	if ent.IsConstraintError(err) {
		tenant2, _ = client.Tenant.Get(ctx, uuid.MustParse("aaaa2222-2222-2222-2222-222222222222"))
	}

	// ✅ Users
	users := []struct {
		ID         uuid.UUID
		Email      string
		Password   string
		Name       string
		IsPlatform bool
		IsTenant   bool
	}{
		{uuid.MustParse("11111111-1111-1111-1111-111111111111"), "alice@example.com", "password123", "Alice", true, true},
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), "bob@example.com", "password123", "Bob", true, false},
		{uuid.MustParse("33333333-3333-3333-3333-333333333333"), "charlie@example.com", "password123", "Charlie", false, true},
	}

	for _, u := range users {
		// Ensure user
		userEntity, err := client.User.Query().Where(user.EmailEQ(u.Email)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return fmt.Errorf("failed to check existing user: %w", err)
		}
		if userEntity == nil {
			hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			userEntity, err = client.User.Create().
				SetID(u.ID).
				SetEmail(u.Email).
				SetPassword(string(hashed)).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create user %s: %w", u.Email, err)
			}
			_, _ = client.Profile.Create().
				SetName(u.Name).
				SetUser(userEntity).
				Save(ctx)
		}

		// Platform user
		if u.IsPlatform {
			exists, _ := client.PlatformUser.Query().
				Where(platformuser.HasUserWith(user.IDEQ(userEntity.ID))).
				Exist(ctx)
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

		// Tenant user
		if u.IsTenant {
			exists, _ := client.TenantUser.Query().
				Where(tenantuser.UserIDEQ(userEntity.ID)).
				Exist(ctx)
			if !exists {
				switch u.Email {
				case "alice@example.com":
					_, err = client.TenantUser.Create().
						SetTenant(tenant1).
						SetUser(userEntity).
						SetStatus(tenantuser.StatusActive).
						Save(ctx)
					if err != nil {
						return fmt.Errorf("failed to create tenant user for Alice: %w", err)
					}
				case "charlie@example.com":
					_, err = client.TenantUser.Create().
						SetTenant(tenant2).
						SetUser(userEntity).
						SetStatus(tenantuser.StatusActive).
						Save(ctx)
					if err != nil {
						return fmt.Errorf("failed to create tenant user for Charlie: %w", err)
					}
				}
			}
		}
	}

	return nil
}
