package services

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/category"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	"github.com/umardev500/go-laundry/internal/types"

	servicesEnt "github.com/umardev500/go-laundry/ent/services"
	appContext "github.com/umardev500/go-laundry/internal/app/context"
	domain "github.com/umardev500/go-laundry/internal/domain/services"
)

type repositoryImpl struct {
	client *db.Client
}

var _ domain.Repository = (*repositoryImpl)(nil)

func NewRepositoryImpl(client *db.Client) domain.Repository {
	return &repositoryImpl{client: client}
}

// Create a new service
func (r *repositoryImpl) Create(ctx *appContext.ScopedContext, payload *domain.Create) (*domain.Services, error) {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	svcEnt, err := conn.Services.
		Create().
		SetTenantID(*scoped.TenantID).
		SetNillableCategoryID(payload.CategoryID).
		SetName(payload.Name).
		SetBasePrice(payload.BasePrice).
		SetNillableDescription(payload.Description).
		SetNillableUnitID(payload.UnitID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(svcEnt), nil
}

func (r *repositoryImpl) GetByID(ctx *appContext.ScopedContext, id uuid.UUID) (*domain.Services, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Services.Query().Where(servicesEnt.IDEQ(id))
	if ctx.Scoped.TenantID != nil {
		q = q.Where(servicesEnt.TenantIDEQ(*ctx.Scoped.TenantID))
	}

	svcEnt, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(svcEnt), nil
}

// List services with filter
func (r *repositoryImpl) List(ctx *appContext.ScopedContext, filter *domain.Filter) (*types.PageData[domain.Services], error) {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	// apply base filters (includes ordering now)
	q := r.applyFilter(conn.Services.Query(), filter, scoped)

	// count before pagination
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// pagination
	q = q.Limit(filter.Limit).Offset(filter.Offset)

	entsList, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Services, len(entsList))
	for i, e := range entsList {
		result[i] = mapFromEnt(e)
	}

	return &types.PageData[domain.Services]{
		Data:  result,
		Total: total,
	}, nil
}

// Update service by ID
func (r *repositoryImpl) Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *domain.Update) (*domain.Services, error) {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	q := conn.Services.UpdateOneID(id).
		SetNillableName(payload.Name).
		SetNillableDescription(payload.Description).
		SetNillableBasePrice(payload.BasePrice).
		SetNillableUnitID(payload.UnitID)

	if scoped.TenantID != nil {
		q = q.Where(servicesEnt.TenantIDEQ(*scoped.TenantID))
	}

	svcEnt, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(svcEnt), nil
}

func (r *repositoryImpl) Delete(ctx *appContext.ScopedContext, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	scoped := ctx.Scoped

	q := conn.Services.DeleteOneID(id)
	if scoped.TenantID != nil {
		q = q.Where(servicesEnt.TenantIDEQ(*scoped.TenantID))
	}

	return q.Exec(ctx)
}

// --- helper methods ---

func (r *repositoryImpl) applyFilter(q *ent.ServicesQuery, filter *domain.Filter, scoped *appContext.Scoped) *ent.ServicesQuery {
	// tenant scoping
	if scoped.TenantID != nil {
		q = q.Where(servicesEnt.TenantIDEQ(*scoped.TenantID))
	}

	// search query
	if filter.Query != "" {
		q = q.Where(servicesEnt.NameContainsFold(filter.Query))
	}

	// include category
	if filter.IncludeCategory {
		q = q.WithCategory()
	}

	// include tenant
	if filter.IncludeTenant {
		q = q.WithTenant()
	}

	// ordering
	switch filter.OrderBy {
	case domain.OrderByNameAsc:
		q = q.Order(ent.Asc(servicesEnt.FieldName))
	case domain.OrderByNameDesc:
		q = q.Order(ent.Desc(servicesEnt.FieldName))
	case domain.OrderByCreatedAtAsc:
		q = q.Order(ent.Asc(servicesEnt.FieldCreatedAt))
	case domain.OrderByCreatedAtDesc:
		q = q.Order(ent.Desc(servicesEnt.FieldCreatedAt))
	default:
		q = q.Order(ent.Asc(servicesEnt.FieldName))
	}

	return q
}

// Mapper from ent -> domain
func mapFromEnt(e *ent.Services) *domain.Services {
	d := &domain.Services{
		ID:          e.ID,
		TenantID:    e.TenantID,
		CategoryID:  e.CategoryID,
		Name:        *e.Name,
		Description: e.Description,
		BasePrice:   *e.BasePrice,
		UnitID:      e.UnitID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}

	// map tenant edge if loaded
	if e.Edges.Tenant != nil {
		d.Tenant = &tenant.Tenant{
			ID:        e.Edges.Tenant.ID,
			Name:      *e.Edges.Tenant.Name,
			Phone:     *e.Edges.Tenant.Phone,
			Email:     *e.Edges.Tenant.Email,
			Address:   *e.Edges.Tenant.Address,
			CreatedAt: e.Edges.Tenant.CreatedAt,
			UpdatedAt: e.Edges.Tenant.UpdatedAt,
		}
	}

	// map category edge if loaded
	if e.Edges.Category != nil {
		d.Category = &category.Category{
			ID:          e.Edges.Category.ID,
			TenantID:    e.Edges.Category.TenantID,
			Name:        *e.Edges.Category.Name,
			Description: e.Edges.Category.Description,
			CreatedAt:   e.Edges.Category.CreatedAt,
			UpdatedAt:   e.Edges.Category.UpdatedAt,
		}
	}

	return d
}
