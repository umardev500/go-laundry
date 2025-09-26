package dto

import "github.com/google/uuid"

type SetPermissionsRequest struct {
	PermissionIDs []uuid.UUID `json:"permission_ids" validate:"required"`
}
