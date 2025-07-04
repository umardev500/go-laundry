package seeds

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/constants"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/feature"
	"github.com/umardev500/go-laundry/internal/ent/permission"
)

var UserManagement = domain.CreateFeatureInput{
	Name:        "User Management",
	Description: "Manage users",
	Enabled:     true,
	Permissions: []domain.CreatePermissionInput{
		{
			Name:        constants.PermissionUserRead.Value,
			Description: constants.PermissionRoleRead.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionUserCreate.Value,
			Description: constants.PermissionUserCreate.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionUserUpdate.Value,
			Description: constants.PermissionUserUpdate.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionUserDelete.Value,
			Description: constants.PermissionUserDelete.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionUserRoleUpdate.Value,
			Description: constants.PermissionUserRoleUpdate.Description,
			Enabled:     true,
		},
	},
}

var RoleManagement = domain.CreateFeatureInput{
	Name:        "Role Management",
	Description: "Manage roles",
	Enabled:     true,
	Permissions: []domain.CreatePermissionInput{
		{
			Name:        constants.PermissionRoleRead.Value,
			Description: constants.PermissionRoleRead.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionRoleCreate.Value,
			Description: constants.PermissionRoleCreate.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionRoleUpdate.Value,
			Description: constants.PermissionRoleUpdate.Description,
			Enabled:     true,
		},
		{
			Name:        constants.PermissionRoleDelete.Value,
			Description: constants.PermissionRoleDelete.Description,
			Enabled:     true,
		},
	},
}

var FeaturesSeed = []domain.CreateFeatureInput{
	UserManagement,
	RoleManagement,
}

func SeedFeatures(ctx context.Context, tx *ent.Tx) ([]uuid.UUID, error) {
	featureIDs := []uuid.UUID{}

	for _, f := range FeaturesSeed {
		// Upsert feature
		ftID, err := tx.Feature.
			Create().
			SetName(f.Name).
			SetDescription(f.Description).
			SetEnabled(f.Enabled).
			OnConflictColumns(feature.FieldName).
			UpdateNewValues().
			ID(ctx)

		// Add feature id to list
		featureIDs = append(featureIDs, ftID)

		if err != nil {
			return featureIDs, err
		}

		// Create permissions
		for _, p := range f.Permissions {
			tx.Permission.
				Create().
				SetName(p.Name).
				SetDescription(p.Description).
				SetEnabled(p.Enabled).
				SetFeatureID(ftID).
				OnConflictColumns(permission.FieldName, permission.FeatureColumn).
				UpdateNewValues().
				ID(ctx)
		}

	}

	return featureIDs, nil
}
