package category

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	categoryEnt "github.com/umardev500/go-laundry/ent/category"
	"github.com/umardev500/go-laundry/internal/db"
	domain "github.com/umardev500/go-laundry/internal/domain/category"
	"github.com/umardev500/go-laundry/internal/types"
)

type repositoryImpl struct {
	client *db.Client
}

// Ensure we implement the domain interface
var _ domain.Repository = (*repositoryImpl)(nil)

func NewRepositoryImpl(client *db.Client) domain.Repository {
	return &repositoryImpl{client: client}
}

func (r *repositoryImpl) Create(ctx context.Context, payload *domain.Create) (*domain.Category, error) {
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

func (r *repositoryImpl) GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) (*domain.Category, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Category.Query().Where(categoryEnt.IDEQ(id))
	if tenantID != nil {
		q = q.Where(categoryEnt.TenantIDEQ(*tenantID))
	}

	catEnt, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(catEnt), nil
}

func (r *repositoryImpl) List(ctx context.Context, tenantID *uuid.UUID, filter domain.Filter) (*types.PageData[domain.Category], error) {
	conn := r.client.GetConn(ctx)

	q := conn.Category.Query()
	if tenantID != nil {
		q = q.Where(categoryEnt.TenantIDEQ(*tenantID))
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

func (r *repositoryImpl) Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *domain.Update) (*domain.Category, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Category.UpdateOneID(id)
	if tenantID != nil {
		exists, err := conn.Category.Query().Where(categoryEnt.IDEQ(id), categoryEnt.TenantIDEQ(*tenantID)).Exist(ctx)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, fmt.Errorf("category not found for this tenant")
		}
	}

	if payload.Name != nil {
		q.SetNillableName(payload.Name)
	}
	if payload.Description != nil {
		q.SetNillableDescription(payload.Description)
	}

	catEnt, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(catEnt), nil
}

func (r *repositoryImpl) Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	q := conn.Category.DeleteOneID(id)
	if tenantID != nil {
		q.Where(categoryEnt.TenantIDEQ(*tenantID))
	}

	return q.Exec(ctx)
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
