package role

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	roleEntity "github.com/umardev500/go-laundry/ent/role"
	"github.com/umardev500/go-laundry/ent/tenant"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/role"
)

type repositoryImpl struct {
	client *db.Client
}

// AssignRoleToUser implements role.Repository.
func (r *repositoryImpl) AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	// Fetch the role
	roleEntity, err := conn.Role.Get(ctx, roleID)
	if err != nil {
		return err
	}

	// Update the user -> attach role
	return conn.User.
		UpdateOneID(userID).
		AddRole(roleEntity).
		Exec(ctx)
}

// Create implements role.Repository.
func (r *repositoryImpl) Create(ctx context.Context, payload *role.RoleCreate, tenantID *uuid.UUID) (*role.Role, error) {
	conn := r.client.GetConn(ctx)

	entRole, err := conn.Role.
		Create().
		SetName(payload.Name).
		SetNillableTenantID(tenantID).
		SetNillableDescription(payload.Description).
		Save(ctx)
	if err != nil {
		fmt.Println("failed:", err)
		return nil, err
	}

	var result role.Role
	r.mapFromEnt(entRole, &result)
	return &result, nil
}

// FindByName implements role.Repository.
func (r *repositoryImpl) FindByName(ctx context.Context, name string, tenantID *uuid.UUID) (*role.Role, error) {
	conn := r.client.GetConn(ctx)

	query := conn.Role.Query()
	if tenantID != nil {
		query = query.Where(roleEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	} else {
		query = query.Where(roleEntity.TenantIDIsNil())
	}

	entRole, err := query.
		Where(roleEntity.NameEQ(name)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	var result role.Role
	r.mapFromEnt(entRole, &result)
	return &result, nil
}

// List implements role.Repository.
func (r *repositoryImpl) List(ctx context.Context, tenantID *uuid.UUID) ([]*role.Role, error) {
	conn := r.client.GetConn(ctx)

	query := conn.Role.Query()
	if tenantID != nil {
		query = query.Where(roleEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	} else {
		query = query.Where(roleEntity.TenantIDIsNil())
	}

	entRoles, err := query.
		All(ctx)

	if err != nil {
		return nil, err
	}

	roles := make([]*role.Role, len(entRoles))
	for i, entRole := range entRoles {
		dest := &role.Role{}
		r.mapFromEnt(entRole, dest)

		roles[i] = dest
	}

	return roles, nil
}

func (r *repositoryImpl) mapFromEnt(e *ent.Role, to *role.Role) {
	if e == nil || to == nil {
		return
	}

	to.ID = e.ID
	to.Name = *e.Name
	to.Description = e.Description
	to.CreatedAt = e.CreatedAt
	to.UpdatedAt = e.UpdatedAt
}

func NewRepository(client *db.Client) role.Repository {
	return &repositoryImpl{
		client: client,
	}
}
