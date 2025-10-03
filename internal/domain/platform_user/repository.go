package platformuser

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
)

// Repository interface
type Repository interface {
	Create(ctx context.Context, payload *Create) (*PlatformUser, error)
	GetByID(ctx context.Context, id uuid.UUID) (*PlatformUser, error)
	GetByUserID(ctx context.Context, id uuid.UUID) (*PlatformUser, error)
	List(ctx context.Context, filter Filter) (*types.PageData[PlatformUser], error)
	Update(ctx context.Context, id uuid.UUID, payload *Update) (*PlatformUser, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
