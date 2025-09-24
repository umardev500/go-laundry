package role

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RoleCreate struct {
	Name        string
	Description *string
}

type RoleUpdate struct {
	Name        *string
	Description *string
}
