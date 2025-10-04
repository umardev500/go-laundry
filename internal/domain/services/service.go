package services

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Service interface {
	Create(ctx *appContext.ScopedContext, payload *Create) (*Services, error)
	GetByID(ctx *appContext.ScopedContext, id uuid.UUID) (*Services, error)
	List(ctx *appContext.ScopedContext, filter *Filter) (*types.PageResult[Services], error)
	Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *Update) (*Services, error)
	Delete(ctx *appContext.ScopedContext, id uuid.UUID) error
}
