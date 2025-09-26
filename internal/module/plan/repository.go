package plan

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	permissionEntity "github.com/umardev500/go-laundry/ent/permission"
	planEntity "github.com/umardev500/go-laundry/ent/plan"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/plan"
)

type repositoryImpl struct {
	client *db.Client
}

// ReplacePermissions implements plan.Repository.
func (r *repositoryImpl) ReplacePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	perms, err := conn.Permission.
		Query().
		Where(permissionEntity.IDIn(permissionIDs...)).
		All(ctx)

	if err != nil {
		return fmt.Errorf("failed to load permissions: %w", err)
	}

	err = conn.Plan.
		UpdateOneID(planID).
		ClearPermissions().
		AddPermissions(perms...).
		Exec(ctx)

	return err
}

// RemovePermissions implements plan.Repository.
func (r *repositoryImpl) RemovePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	perms, err := conn.Permission.
		Query().
		Where(permissionEntity.IDIn(permissionIDs...)).
		All(ctx)

	if err != nil {
		return fmt.Errorf("failed to load permissions: %w", err)
	}

	err = conn.Plan.
		UpdateOneID(planID).
		RemovePermissions(perms...).
		Exec(ctx)

	return err
}

// AddPermissions implements plan.Repository.
func (r *repositoryImpl) AddPermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	plantEnt, err := conn.Plan.Get(ctx, planID)
	if err != nil {
		return fmt.Errorf("failed to load plan: %w", err)
	}

	perms, err := conn.Permission.
		Query().
		Where(permissionEntity.IDIn(permissionIDs...)).
		All(ctx)

	if err != nil {
		return fmt.Errorf("failed to load permissions: %w", err)
	}

	err = conn.Plan.
		UpdateOne(plantEnt).
		AddPermissions(perms...).
		Exec(ctx)

	return err
}

// GetByID implements plan.Repository.
func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*plan.Plan, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Plan.Query()
	entPlan, err := q.Where(planEntity.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entPlan), nil
}

// List implements plan.Repository.
func (r *repositoryImpl) List(ctx context.Context, filter *plan.PlanFilter) ([]*plan.Plan, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Plan.Query()

	if !filter.IncludeDeleted {
		q = q.Where(planEntity.DeletedAtIsNil())
	}

	plansEnt, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnts(plansEnt), nil
}

func (r *repositoryImpl) mapFromEnts(es []*ent.Plan) []*plan.Plan {
	var result []*plan.Plan
	for _, e := range es {
		result = append(result, r.mapFromEnt(e))
	}
	return result
}

func (r *repositoryImpl) mapFromEnt(e *ent.Plan) *plan.Plan {
	return &plan.Plan{
		ID:          e.ID,
		Name:        *e.Name,
		Description: e.Description,
		MaxOrders:   e.MaxOrders,
		MaxUsers:    e.MaxUsers,
		Price:       e.Price,
		Duration:    e.DurationDays,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func NewRepository(client *db.Client) plan.Repository {
	return &repositoryImpl{
		client: client,
	}
}
