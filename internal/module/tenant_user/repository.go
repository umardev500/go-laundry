package tenantuser

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	tenantuserEnt "github.com/umardev500/go-laundry/ent/tenantuser"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	domain "github.com/umardev500/go-laundry/internal/domain/tenant_user"
	"github.com/umardev500/go-laundry/internal/types"
)

type repositoryImpl struct {
	client *db.Client
}

var _ domain.Repository = (*repositoryImpl)(nil)

func NewRepositoryImpl(client *db.Client) domain.Repository {
	return &repositoryImpl{client: client}
}

func (r *repositoryImpl) Create(ctx context.Context, payload *domain.Create) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTU, err := conn.TenantUser.
		Create().
		SetTenantID(payload.TenantID).
		SetUserID(payload.UserID).
		SetNillableStatus((*tenantuserEnt.Status)(payload.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(entTU), nil
}

func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTU, err := conn.TenantUser.Query().
		Where(tenantuserEnt.IDEQ(id)).
		WithTenant().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("tenant user not found: %w", err)
	}

	return mapFromEnt(entTU), nil
}

func (r *repositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID, filter *domain.Filter) (*types.PageData[domain.TenantUser], error) {
	conn := r.client.GetConn(ctx)

	q := conn.TenantUser.Query().
		Where(tenantuserEnt.UserIDEQ(userID)).
		WithTenant()

	// apply ordering
	q = applyFilter(q, filter)

	// count total
	total, err := q.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count tenant users by userID: %w", err)
	}

	// paging
	q = q.Limit(filter.Limit).Offset(filter.Offset)

	// fetch
	entList, err := q.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant users by userID: %w", err)
	}

	// map ent -> domain
	result := make([]*domain.TenantUser, len(entList))
	for i, e := range entList {
		result[i] = mapFromEnt(e)
	}

	return &types.PageData[domain.TenantUser]{
		Data:  result,
		Total: total,
	}, nil
}

func (r *repositoryImpl) List(ctx context.Context, filter *domain.Filter) (*types.PageData[domain.TenantUser], error) {
	conn := r.client.GetConn(ctx)

	q := conn.TenantUser.
		Query().
		WithTenant()

	q = applyFilter(q, filter)

	total, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	q = q.Limit(filter.Limit).Offset(filter.Offset)

	entList, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.TenantUser, len(entList))
	for i, e := range entList {
		result[i] = mapFromEnt(e)
	}

	return &types.PageData[domain.TenantUser]{Data: result, Total: total}, nil
}

func (r *repositoryImpl) Update(ctx context.Context, id uuid.UUID, payload *domain.Update) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	q := conn.TenantUser.UpdateOneID(id).
		SetNillableStatus((*tenantuserEnt.Status)(payload.Status))

	entTU, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(entTU), nil
}

func (r *repositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	return conn.TenantUser.DeleteOneID(id).Exec(ctx)
}

func applyFilter(q *ent.TenantUserQuery, filter *domain.Filter) *ent.TenantUserQuery {
	switch filter.OrderBy {
	case domain.OrderByCreatedAtDesc:
		q = q.Order(ent.Desc(tenantuserEnt.FieldCreatedAt))
	case domain.OrderByUpdatedAtAsc:
		q = q.Order(ent.Asc(tenantuserEnt.FieldUpdatedAt))
	case domain.OrderByUpdatedAtDesc:
		q = q.Order(ent.Desc(tenantuserEnt.FieldUpdatedAt))
	default:
		q = q.Order(ent.Asc(tenantuserEnt.FieldCreatedAt))
	}
	return q
}

func mapFromEnt(e *ent.TenantUser) *domain.TenantUser {
	var mappedTenant *tenant.Tenant
	tenantEnt := e.Edges.Tenant
	if tenantEnt != nil {
		mappedTenant = &tenant.Tenant{
			ID:        tenantEnt.ID,
			Name:      *tenantEnt.Name,
			Phone:     *tenantEnt.Phone,
			Email:     *tenantEnt.Email,
			Address:   *tenantEnt.Address,
			CreatedAt: tenantEnt.CreatedAt,
			UpdatedAt: tenantEnt.UpdatedAt,
		}
	}

	return &domain.TenantUser{
		ID:        e.ID,
		TenantID:  *e.TenantID,
		Tenant:    mappedTenant,
		UserID:    *e.UserID,
		Status:    domain.Status(*e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: nil, // soft deletes later if needed
	}
}
