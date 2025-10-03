package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
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

type UserUpdate struct {
	Email          *string
	Password       *string
	ResetToken     *string
	ResetExpiresAt *time.Time
}

type UserOrderBy string

const (
	OrderByEmailAsc      UserOrderBy = "email_asc"
	OrderByEmailDesc     UserOrderBy = "email_desc"
	OrderByCreatedAtAsc  UserOrderBy = "created_at_asc"
	OrderByCreatedAtDesc UserOrderBy = "created_at_desc"
)

type Filter struct {
	TenantID *uuid.UUID  `query:"-"`
	Scope    types.Scope `query:"-"`

	Query          string      `query:"query"` // search by email or name
	Limit          int         `query:"limit"` // pagination
	Offset         int         `query:"offset"`
	OrderBy        UserOrderBy `query:"order_by"`        // e.g. "email asc", "created_at desc"
	IncludeDeleted bool        `query:"include_deleted"` // if true, include soft-deleted users
}

func (f *Filter) WithDefaults() *Filter {
	if f.Limit == 0 {
		f.Limit = 10 // default page size
	}
	if f.Offset == 0 {
		f.Offset = 0
	}
	if f.OrderBy == "" {
		f.OrderBy = "created_at desc" // default ordering
	}
	return f
}

func (f *Filter) Validate() error {
	switch f.Scope {
	case types.ScopeTenant:
		if f.TenantID == nil {
			return fmt.Errorf("tenant id is required")
		}
	case types.ScopePlatform:
		if f.TenantID != nil {
			return fmt.Errorf("tenant id is not allowed")
		}
	case types.ScopeGlobal:
		if f.TenantID != nil {
			return fmt.Errorf("tenant id is not allowed")
		}
	default:
		return fmt.Errorf("invalid scope: %s", f.Scope)
	}

	return nil
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
	Name   string
	Avatar *string
	Phone  *string
}

type ProfileUpdate struct {
	Name   *string
	Avatar *string
	Phone  *string
}
