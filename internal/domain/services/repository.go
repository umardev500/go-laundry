package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
)

type Repository interface {
	Create(ctx context.Context, payload *Create, tenantID *uuid.UUID) (*Services, error)
	GetByID(ctx context.Context, id uuid.UUID, tenantID *uuid.UUID) (*Services, error)
	List(ctx context.Context, filter *Filter, tenantID *uuid.UUID) (*types.PageData[Services], error)
	Update(ctx context.Context, id uuid.UUID, payload *Update, tenantID *uuid.UUID) (*Services, error)
	Delete(ctx context.Context, id uuid.UUID, tenantID *uuid.UUID) error
}
