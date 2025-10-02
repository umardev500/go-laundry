package category

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, payload *Create) (*Category, error)
	GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) (*Category, error)
	List(ctx context.Context, tenantID *uuid.UUID, filter Filter) ([]*Category, error)
	Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *Update) (*Category, error)
	Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error
}
