package errors

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
)

var (
	ErrTenantNotFound = fmt.Errorf("tenant not found")
)

func NewTenantNotFound(id uuid.UUID) *core.Error {
	return core.NewError(
		ErrTenantNotFound,
		fmt.Sprintf("Tenant with id %s not found", id.String()),
		http.StatusNotFound,
	)
}
