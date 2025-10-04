package registration

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/permission"
	"github.com/umardev500/go-laundry/internal/domain/registration"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	tenantuser "github.com/umardev500/go-laundry/internal/domain/tenant_user"
	"github.com/umardev500/go-laundry/internal/domain/user"
)

type service struct {
	userService       user.Service
	tenantService     tenant.Service
	roleService       role.Service
	tenantUserService tenantuser.Service
	permissionService permission.Service
	client            *db.Client
}

// RegisterUser implements registration.Service.
func (s *service) RegisterUser(ctx context.Context, payload *registration.CreateUser) (*user.User, error) {
	var usr *user.User

	err := s.client.WithTransaction(ctx, func(ctx context.Context) error {
		var err error

		// Create user
		usr, err = s.userService.Create(ctx, payload.User)
		if err != nil {
			return err
		}

		// Create user profile
		_, err = s.userService.CreateProfile(ctx, usr.ID, payload.Profile)
		if err != nil {
			return err
		}

		// -- Assign role ---
		role, err := s.roleService.GetRoleByName(ctx, "customer", nil)
		if err != nil {
			return err
		}

		err = s.roleService.AssignRoleToUser(ctx, nil, usr.ID, role.ID)
		if err != nil {
			return err
		}

		// Assign default permissions to user
		defaultPermissions, err := s.getDefaultPermissions(ctx)
		if err != nil {
			return err
		}
		err = s.permissionService.AssignPermissionsToRole(ctx, role.ID, defaultPermissions)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// Ensure serviceImpl implements the domain Service interface
var _ registration.Service = (*service)(nil)

func NewService(
	userService user.Service,
	tenantService tenant.Service,
	tenantUserService tenantuser.Service,
	roleService role.Service,
	permissionService permission.Service,
	client *db.Client,
) *service {
	return &service{
		userService:       userService,
		tenantService:     tenantService,
		tenantUserService: tenantUserService,
		roleService:       roleService,
		permissionService: permissionService,
		client:            client,
	}
}

func (s *service) RegisterTenant(ctx context.Context, data *registration.CreateTenantUser) (tnt *tenant.Tenant, err error) {
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

		tenantID := func() *uuid.UUID {
			id := t.ID
			return &id
		}()

		data.User.TenantID = tenantID

		var usr *user.User

		// Create default tenant role
		tenantRole, err := s.roleService.CreateRole(ctx, &role.RoleCreate{
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

		// Try to find existing user
		existingUser, err := s.userService.FindByEmail(ctx, data.User.Email)
		if err == nil {
			usr = existingUser
		} else if ent.IsNotFound(err) {
			// Not found -> create new user
			usr, err = s.userService.Create(ctx, data.User)
			if err != nil {
				return err
			}
		} else {
			// Unexpected error
			return fmt.Errorf("failed to check user existence: %w", err)
		}

		// Create user profile
		_, err = s.userService.CreateProfile(ctx, usr.ID, data.Profile)
		if err != nil {
			return err
		}

		// Create tenant user
		_, err = s.tenantUserService.Create(ctx, &tenantuser.Create{
			UserID:   usr.ID,
			TenantID: *tenantID,
		})
		if err != nil {
			return err
		}

		// Assign role to user
		err = s.roleService.AssignRoleToUser(ctx, tenantID, usr.ID, tenantRole.ID)
		if err != nil {
			fmt.Println(usr.ID)
			log.Error().Err(err).Msg("failed to create user profile")
			return err
		}

		// Assign default permissions to user
		err = s.permissionService.AssignPermissionsToRole(ctx, tenantRole.ID, defaultPermissions)
		if err != nil {
			return err
		}

		tnt = t

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
