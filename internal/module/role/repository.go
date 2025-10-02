package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	roleEntity "github.com/umardev500/go-laundry/ent/role"
	"github.com/umardev500/go-laundry/ent/tenant"
	userEntity "github.com/umardev500/go-laundry/ent/user"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/permission"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/types"
)

type repositoryImpl struct {
	client *db.Client
}

// AssignRoleToUser implements role.Repository.
func (r *repositoryImpl) AssignRoleToUser(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	userQuery := conn.User.
		Query().
		Where(userEntity.IDEQ(userID))

	roleQuery := conn.Role.
		Query().
		Where(roleEntity.IDEQ(roleID))

	// If tenantID is not nil, enforce tenant scoping
	if tenantID != nil {
		userQuery = userQuery.
			Where(userEntity.TenantIDEQ(*tenantID))

		roleQuery = roleQuery.
			Where(roleEntity.TenantIDEQ(*tenantID))
	}

	userEnt, err := userQuery.Only(ctx)
	if err != nil {
		return err
	}

	roleEnt, err := roleQuery.Only(ctx)
	if err != nil {
		return err
	}

	// Update the user -> attach role
	return conn.User.
		UpdateOne(userEnt).
		AddRole(roleEnt).
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
func (r *repositoryImpl) List(ctx context.Context, f *role.Filter, tenantID *uuid.UUID) (*types.PageData[role.Role], error) {
	conn := r.client.GetConn(ctx)

	q := conn.Role.Query()
	if tenantID != nil {
		q = q.Where(roleEntity.HasTenantWith(tenant.IDEQ(*tenantID)))
	} else {
		q = q.Where(roleEntity.TenantIDIsNil())
	}

	// --- Apply filters ---
	if !f.IncludeDeleted {
		q = q.Where(roleEntity.DeletedAtIsNil())
	}

	if f.IncludePermissions {
		q = q.WithPermissions()
	}

	if f.Query != "" {
		q = q.Where(roleEntity.NameContainsFold(f.Query))
	}

	// Count total
	total, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	// Ordering
	switch f.OrderBy {
	case role.OrderByCreatedAtDesc:
		q = q.Order(ent.Desc(roleEntity.FieldCreatedAt))
	default:
		q = q.Order(ent.Asc(roleEntity.FieldCreatedAt))
	}

	// Pagination
	q = q.Limit(f.Limit).Offset(f.Offset)

	entRoles, err := q.
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

	return &types.PageData[role.Role]{
		Data:  roles,
		Total: total,
	}, nil
}

func (r *repositoryImpl) mapFromEnt(e *ent.Role, to *role.Role) {
	if e == nil || to == nil {
		return
	}

	var mappedPermissions []*permission.Permission
	if e.Edges.Permissions != nil {
		for _, p := range e.Edges.Permissions {
			mappedPermissions = append(mappedPermissions, &permission.Permission{
				ID:   p.ID,
				Name: *p.Name,
			})
		}
	}

	to.ID = e.ID
	to.Name = *e.Name
	to.Description = e.Description
	to.Permissions = mappedPermissions
	to.CreatedAt = e.CreatedAt
	to.UpdatedAt = e.UpdatedAt
}

func NewRepository(client *db.Client) role.Repository {
	return &repositoryImpl{
		client: client,
	}
}
