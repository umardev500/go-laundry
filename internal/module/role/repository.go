package role

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/ent/tenant"
	"github.com/umardev500/go-laundry/ent/tenantuser"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/permission"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/types"

	roleEntity "github.com/umardev500/go-laundry/ent/role"
	userEntity "github.com/umardev500/go-laundry/ent/user"
	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type repositoryImpl struct {
	client *db.Client
}

// AssignRoleToUser implements role.Repository.
func (r *repositoryImpl) AssignRoleToUser(ctx *appContext.ScopedContext, userID uuid.UUID, roleID uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped
	tenantID := scoped.TenantID

	userQuery := conn.User.
		Query().
		Where(userEntity.IDEQ(userID))

	roleQuery := conn.Role.
		Query().
		Where(roleEntity.IDEQ(roleID))

	// If tenantID is not nil, enforce tenant scoping
	if tenantID != nil {
		userQuery = userQuery.
			Where(userEntity.HasTenantUsersWith(tenantuser.TenantIDEQ(*tenantID)))

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
func (r *repositoryImpl) Create(ctx *appContext.ScopedContext, payload *role.RoleCreate) (*role.Role, error) {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	entRole, err := conn.Role.
		Create().
		SetName(payload.Name).
		SetNillableTenantID(scoped.TenantID).
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
func (r *repositoryImpl) FindByName(ctx *appContext.ScopedContext, name string) (*role.Role, error) {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	query := conn.Role.Query()
	if scoped.TenantID != nil {
		query = query.Where(roleEntity.HasTenantWith(tenant.IDEQ(*scoped.TenantID)))
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
func (r *repositoryImpl) List(ctx *appContext.ScopedContext, f *role.Filter) (*types.PageData[role.Role], error) {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped
	tenantID := scoped.TenantID

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
