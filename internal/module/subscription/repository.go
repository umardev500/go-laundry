package subscription

import (
	"context"

	"github.com/umardev500/go-laundry/ent"
	subscriptionEntity "github.com/umardev500/go-laundry/ent/subscription"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
)

type repositoryImpl struct {
	client *db.Client
}

// Create implements subscription.Repository.
func (r *repositoryImpl) Create(ctx context.Context, payload *subscription.SubscriptionCreate) (*subscription.Subscription, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.Subscription.
		Create().
		SetPlanID(payload.PlanID).
		SetTenantID(payload.TenantID).
		SetNillableStartDate(payload.StartDate).
		SetNillableEndDate(payload.EndDate).
		SetNillableStatus((*subscriptionEntity.Status)(payload.Status))

	subEnt, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(subEnt), nil
}

// List implements subscription.Repository.
func (r *repositoryImpl) List(ctx context.Context) ([]*subscription.Subscription, error) {
	conn := r.client.GetConn(ctx)

	subs, err := conn.Subscription.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnts(subs), nil
}

func (r *repositoryImpl) mapFromEnts(es []*ent.Subscription) []*subscription.Subscription {
	var result []*subscription.Subscription
	for _, e := range es {
		result = append(result, r.mapFromEnt(e))
	}
	return result
}

func (r *repositoryImpl) mapFromEnt(s *ent.Subscription) *subscription.Subscription {
	return &subscription.Subscription{
		ID:        s.ID,
		PlanID:    s.PlanID,
		TenantID:  s.TenantID,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		Status:    subscription.SubscriptionStatus(s.Status),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func NewRepository(client *db.Client) subscription.Repository {
	return &repositoryImpl{
		client: client,
	}
}
