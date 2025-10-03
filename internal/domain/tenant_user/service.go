package tenantuser

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
)

type Service interface {
	Create(ctx context.Context, payload *Create) (*TenantUser, error)
	GetByID(ctx context.Context, id uuid.UUID) (*TenantUser, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*TenantUser, error)
	List(ctx context.Context, filter Filter) (*types.PageResult[TenantUser], error)
	Update(ctx context.Context, id uuid.UUID, payload *Update) (*TenantUser, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
