package plan

import (
	"context"

	"github.com/umardev500/go-laundry/ent"
	planEntity "github.com/umardev500/go-laundry/ent/plan"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/plan"
)

type repositoryImpl struct {
	client *db.Client
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
