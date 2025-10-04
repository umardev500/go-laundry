package context

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/pkg/response"
)

var (
	ErrScopedRequired = fmt.Errorf("scoped is required")
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

type scopedKeyType struct{}

var scopedKey = scopedKeyType{}

const ScopedContextKey = "scopedContext"

// WithScoped attaches a scoped object to the context/
func (s *Scoped) WithScoped(ctx context.Context) context.Context {
	return context.WithValue(ctx, scopedKey, s)
}

func GetScopedContext(c *fiber.Ctx) *ScopedContext {
	val := c.Locals(ScopedContextKey)
	if scopedCtx, ok := val.(*ScopedContext); ok {
		return scopedCtx
	}

	c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
		Success: false,
		Error:   "Scoped context not available or invalid",
	})

	return nil
}

// Validate ensures the scoped instance matches scope rules.
func (s *Scoped) Validate() error {
	switch s.Scope {
	case ScopeTenant:
		if s.TenantID == nil {
			return fmt.Errorf("tenant id is required")
		}
	case ScopePlatform:
		if s.TenantID != nil {
			return fmt.Errorf("tenant id is not allowed")
		}
	case ScopeGlobal:
		if s.TenantID != nil {
			return fmt.Errorf("tenant id is not allowed")
		}
	default:
		return fmt.Errorf("invalid scope: %s", s.Scope)
	}

	return nil
}
