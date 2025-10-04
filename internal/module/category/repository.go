package category

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/ent/predicate"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/types"

	categoryEnt "github.com/umardev500/go-laundry/ent/category"
	appContext "github.com/umardev500/go-laundry/internal/app/context"
	domain "github.com/umardev500/go-laundry/internal/domain/category"
)

type repositoryImpl struct {
	client *db.Client
}

// Ensure we implement the domain interface
var _ domain.Repository = (*repositoryImpl)(nil)

func NewRepositoryImpl(client *db.Client) domain.Repository {
	return &repositoryImpl{client: client}
}

func (r *repositoryImpl) Create(ctx *appContext.ScopedContext, payload *domain.Create) (*domain.Category, error) {
	conn := r.client.GetConn(ctx)

	catEnt, err := conn.Category.
		Create().
		SetNillableTenantID(payload.TenantID).
		SetName(payload.Name).
		SetNillableDescription(payload.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(catEnt), nil
}

func (r *repositoryImpl) GetByID(ctx *appContext.ScopedContext, id uuid.UUID) (*domain.Category, error) {
	var err error
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	q := conn.Category.Query().Where(categoryEnt.IDEQ(id))

	q, err = applyScopeFilter(q, scoped)
	if err != nil {
		return nil, err
	}

	catEnt, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(catEnt), nil
}

func (r *repositoryImpl) List(ctx *appContext.ScopedContext, filter domain.Filter) (*types.PageData[domain.Category], error) {
	var err error

	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	q := conn.Category.Query()
	q, err = applyScopeFilter(q, scoped)
	if err != nil {
		return nil, err
	}

	if filter.Query != "" {
		q = q.Where(categoryEnt.NameContainsFold(filter.Query))
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	switch filter.OrderBy {
	case domain.OrderByNameAsc:
		q = q.Order(ent.Asc(categoryEnt.FieldName))
	case domain.OrderByNameDesc:
		q = q.Order(ent.Desc(categoryEnt.FieldName))
	default:
		q = q.Order(ent.Asc(categoryEnt.FieldName))
	}

	q = q.Limit(filter.Limit).Offset(filter.Offset)

	entsList, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Category, len(entsList))
	for i, e := range entsList {
		result[i] = mapFromEnt(e)
	}

	return &types.PageData[domain.Category]{
		Data:  result,
		Total: total,
	}, nil
}

func (r *repositoryImpl) Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *domain.Update) (*domain.Category, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Category.UpdateOneID(id).
		SetNillableName(payload.Name).
		SetNillableDescription(payload.Description)

	catEnt, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(catEnt), nil
}

func (r *repositoryImpl) Delete(ctx *appContext.ScopedContext, id uuid.UUID) error {
	var err error

	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	q := conn.Category.DeleteOneID(id)
	q, err = applyScopeFilter(q, scoped)
	if err != nil {
		return err
	}

	return q.Exec(ctx)
}

func applyScopeFilter[T interface {
	Where(...predicate.Category) T
}](q T, scoped *appContext.Scoped) (T, error) {
	switch scoped.Scope {
	case appContext.ScopeTenant:
		q = q.Where(
			categoryEnt.TenantIDEQ(*scoped.TenantID),
		)
	}

	return q, nil
}

func mapFromEnt(e *ent.Category) *domain.Category {
	return &domain.Category{
		ID:          e.ID,
		TenantID:    e.TenantID,
		Name:        *e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
