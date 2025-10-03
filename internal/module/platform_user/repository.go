package platformuser

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	platformuserEnt "github.com/umardev500/go-laundry/ent/platformuser"
	"github.com/umardev500/go-laundry/internal/db"
	domain "github.com/umardev500/go-laundry/internal/domain/platform_user"
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

func (r *repositoryImpl) Create(ctx context.Context, payload *domain.Create) (*domain.PlatformUser, error) {
	conn := r.client.GetConn(ctx)

	entUser, err := conn.PlatformUser.
		Create().
		SetUserID(payload.UserID).
		SetNillableStatus((*platformuserEnt.Status)(payload.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(entUser), nil
}

func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	conn := r.client.GetConn(ctx)

	entUser, err := conn.PlatformUser.Query().
		Where(platformuserEnt.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("platform user not found: %w", err)
	}

	return mapFromEnt(entUser), nil
}

func (r *repositoryImpl) GetByUserID(ctx context.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	conn := r.client.GetConn(ctx)

	entUser, err := conn.PlatformUser.Query().
		Where(platformuserEnt.UserIDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("platform user not found: %w", err)
	}

	return mapFromEnt(entUser), nil
}

func (r *repositoryImpl) List(ctx context.Context, filter domain.Filter) (*types.PageData[domain.PlatformUser], error) {
	conn := r.client.GetConn(ctx)

	q := conn.PlatformUser.Query()
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

	result := make([]*domain.PlatformUser, len(entList))
	for i, e := range entList {
		result[i] = mapFromEnt(e)
	}

	return &types.PageData[domain.PlatformUser]{
		Data:  result,
		Total: total,
	}, nil
}

func (r *repositoryImpl) Update(ctx context.Context, id uuid.UUID, payload *domain.Update) (*domain.PlatformUser, error) {
	conn := r.client.GetConn(ctx)

	q := conn.PlatformUser.
		UpdateOneID(id).
		SetNillableStatus((*platformuserEnt.Status)(payload.Status))

	entUser, err := q.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(entUser), nil
}

func (r *repositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	return conn.PlatformUser.DeleteOneID(id).Exec(ctx)
}

// applyFilter applies ordering (no query search for now)
func applyFilter(q *ent.PlatformUserQuery, filter domain.Filter) *ent.PlatformUserQuery {
	switch filter.OrderBy {
	case domain.OrderByCreatedAtDesc:
		q = q.Order(ent.Desc(platformuserEnt.FieldCreatedAt))
	case domain.OrderByUpdatedAtAsc:
		q = q.Order(ent.Asc(platformuserEnt.FieldUpdatedAt))
	case domain.OrderByUpdatedAtDesc:
		q = q.Order(ent.Desc(platformuserEnt.FieldUpdatedAt))
	default:
		q = q.Order(ent.Asc(platformuserEnt.FieldCreatedAt)) // default
	}
	return q
}

func mapFromEnt(e *ent.PlatformUser) *domain.PlatformUser {
	return &domain.PlatformUser{
		ID:        e.ID,
		UserID:    *e.UserID,
		Status:    domain.Status(*e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: nil, // map if using soft deletes
	}
}
