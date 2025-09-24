package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID  `json:"id"`
	TenantID       *uuid.UUID `json:"tenant_id"`
	Email          string     `json:"email"`
	Password       string     `json:"-"`
	ResetToken     *string    `json:"reset_token,omitempty"`
	ResetExpiresAt *time.Time `json:"reset_expires_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type UserCreate struct {
	Email    string
	Password string
	TenantID *uuid.UUID
}

type Profile struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Avatar  *string   `json:"avatar,omitempty"`
	Phone   *string   `json:"phone,omitempty"`
	Address *string   `json:"address,omitempty"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type ProfileCreate struct {
	Name    string
	Avatar  *string
	Phone   *string
	Address *string
}

type ProfileUpdate struct {
	Name    *string
	Avatar  *string
	Phone   *string
	Address *string
}
