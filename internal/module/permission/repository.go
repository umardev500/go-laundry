package permission

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	permissionEntity "github.com/umardev500/go-laundry/ent/permission"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/permission"
)

type repositoryImpl struct {
	client *db.Client
}

// GetPermissionsByNames implements permission.Repository.
func (r *repositoryImpl) GetPermissionsByNames(ctx context.Context, names []string) ([]permission.Permission, error) {
	conn := r.client.GetConn(ctx)
	permissionsEnt, err := conn.Permission.
		Query().
		Where(permissionEntity.NameIn(names...)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	var permissions []permission.Permission
	for _, p := range permissionsEnt {
		var permission permission.Permission
		r.mapFromEnt(p, &permission)
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

// AssignPermissionToRole implements permission.Repository.
func (r *repositoryImpl) AssignPermissionsToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	// Fetch the role
	roleEntity, err := conn.Role.Get(ctx, roleID)
	if err != nil {
		return err
	}

	// Fetch the permission
	permissionEntities, err := conn.Permission.
		Query().
		Where(permissionEntity.IDIn(permissionIDs...)).
		All(ctx)
	if err != nil {
		return err
	}

	// Create the edge between role and permission
	return roleEntity.Update().
		AddPermissions(permissionEntities...).
		Exec(ctx)

}

func (r *repositoryImpl) mapFromEnt(entPermission *ent.Permission, domainPermission *permission.Permission) {
	domainPermission.ID = entPermission.ID
	domainPermission.Name = *entPermission.Name
	domainPermission.CreatedAt = entPermission.CreatedAt
	domainPermission.UpdatedAt = entPermission.UpdatedAt
}

func NewRepository(client *db.Client) permission.Repository {
	return &repositoryImpl{
		client: client,
	}
}
