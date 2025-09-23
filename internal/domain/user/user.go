package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
)

type User struct {
	ID             uuid.UUID  `json:"id"`
	Email          string     `json:"email"`
	Password       string     `json:"-"`
	ResetToken     *string    `json:"reset_token,omitempty"`
	ResetExpiresAt *time.Time `json:"reset_expires_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// MapFromEnt sets the fields of the domain User from an ent.User
func (u *User) MapFromEnt(e *ent.User) {
	if u == nil {
		return
	}

	u.ID = e.ID
	u.Email = e.Email
	u.Password = e.Password
	u.ResetToken = e.ResetToken
	u.ResetExpiresAt = e.ResetExpiresAt
	u.CreatedAt = e.CreatedAt
	u.UpdatedAt = e.UpdatedAt
}

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
}
