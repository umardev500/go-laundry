package types

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Scope string

const (
	ScopePlatform Scope = "platform"
	ScopeTenant   Scope = "tenant"
	ScopeGlobal   Scope = "end"
)

type Scoped struct {
	TenantID *uuid.UUID
	Scope    Scope
}

func (f *Scoped) Validate() error {
	switch f.Scope {
	case ScopeTenant:
		if f.TenantID == nil {
			return fmt.Errorf("tenant id is required")
		}
	case ScopePlatform:
		if f.TenantID != nil {
			return fmt.Errorf("tenant id is not allowed")
		}
	case ScopeGlobal:
		if f.TenantID != nil {
			return fmt.Errorf("tenant id is not allowed")
		}
	default:
		return fmt.Errorf("invalid scope: %s", f.Scope)
	}

	return nil
}

// Seeder is a common interface for all module seeders.
type Seeder interface {
	Seed(ctx context.Context) error
}
