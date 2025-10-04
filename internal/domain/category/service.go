package category

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Service interface {
	Create(ctx *appContext.ScopedContext, payload *Create) (*Category, error)
	GetByID(ctx *appContext.ScopedContext, id uuid.UUID) (*Category, error)
	List(ctx *appContext.ScopedContext, filter Filter) (*types.PageResult[Category], error)
	Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *Update) (*Category, error)
	Delete(ctx *appContext.ScopedContext, id uuid.UUID) error
}
