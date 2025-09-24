package registration

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/permission"
	"github.com/umardev500/go-laundry/internal/domain/registration"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type service struct {
	userService       user.Service
	tenantService     tenant.Service
	roleService       role.Service
	permissionService permission.Service
	client            *db.Client
}

func NewService(
	userService user.Service,
	tenantService tenant.Service,
	roleService role.Service,
	permissionService permission.Service,
	client *db.Client,
) *service {
	return &service{
		userService:       userService,
		tenantService:     tenantService,
		roleService:       roleService,
		permissionService: permissionService,
		client:            client,
	}
}

func (s *service) RegisterTenant(ctx context.Context, data *registration.RegisterInput) (usr *user.User, err error) {
	defaultPermissions, err := s.getDefaultPermissions(ctx)
	if err != nil {
		return nil, err
	}

	err = s.client.WithTransaction(ctx, func(ctx context.Context) error {
		// Create tenant first
		t, err := s.tenantService.CreateTenant(ctx, data.Tenant)
		if err != nil {
			return err
		}

		data.User.TenantID = func() *uuid.UUID {
			id := t.ID
			return &id
		}()

		// Create default tenant user role
		userRole, err := s.roleService.CreateRole(ctx, &role.RoleCreate{
			Name: "admin",
			Description: func() *string {
				desc := "Tenant admin"
				return &desc
			}(),
		}, func() *uuid.UUID {
			id := t.ID
			return &id
		}())
		if err != nil {
			return err
		}

		// Create user
		usr, err = s.userService.CreateUser(ctx, data.User)
		if err != nil {
			return err
		}

		// Create user profile
		_, err = s.userService.CreateProfile(ctx, usr.ID, data.Profile)
		if err != nil {
			return err
		}

		// Assign default permissions to user
		err = s.permissionService.AssignPermissionsToRole(ctx, userRole.ID, defaultPermissions)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to register tenant: %w", err)
	}

	return
}

func (s *service) getDefaultPermissions(ctx context.Context) ([]uuid.UUID, error) {
	perms, err := s.permissionService.GetPermissionsByNames(ctx, []string{
		"create_order",
	})
	if err != nil {
		return nil, err
	}

	var ids []uuid.UUID
	for _, p := range perms {
		ids = append(ids, p.ID)
	}

	return ids, nil
}
