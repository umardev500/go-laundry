package subscription

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	subscriptionEntity "github.com/umardev500/go-laundry/ent/subscription"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/plan"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
	"github.com/umardev500/go-laundry/internal/utils/redisutils"
)

type repositoryImpl struct {
	client      *db.Client
	redisClient *db.RedisClient
}

// GetByID implements subscription.Repository.
func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID, filter *subscription.SubscriptionFilter) (*subscription.Subscription, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Subscription.
		Query().
		Where(subscriptionEntity.IDEQ(id))

	if filter.IncludePlan {
		q = q.WithPlan()
	}

	if filter.IncludeTenant {
		q = q.WithTenant()
	}

	sub, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(sub), nil
}

// Update implements subscription.Repository.
func (r *repositoryImpl) Update(ctx context.Context, payload *subscription.SubscriptionUpdate, id uuid.UUID) (*subscription.Subscription, error) {
	conn := r.client.GetConn(ctx)

	sub, err := conn.Subscription.
		UpdateOneID(id).
		SetNillableStartDate(payload.StartDate).
		SetNillableEndDate(payload.EndDate).
		SetNillableStatus((*subscriptionEntity.Status)(payload.Status)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	// --- Redis cache update for current active plan ---
	//
	// If the subscription is active, store the tnant's current plan in Redis for fast access.
	if payload.Status != nil && *payload.Status == subscription.SubscriptionStatusActive {
		err := r.setActivePlan(ctx, *sub.TenantID, *sub.PlanID, *sub.EndDate)
		if err != nil {
			return nil, err
		}
	}

	return r.mapFromEnt(sub), nil
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

	sub, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	subEnt, err := conn.Subscription.
		Query().
		WithPlan().
		WithTenant().
		Where(subscriptionEntity.IDEQ(sub.ID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// --- Redis cache update for current active plan ---
	//
	// If the subscription is active, store the tnant's current plan in Redis for fast access.
	if payload.Status != nil && *payload.Status == subscription.SubscriptionStatusActive {
		err := r.setActivePlan(ctx, *sub.TenantID, *sub.PlanID, *sub.EndDate)
		if err != nil {
			return nil, err
		}
	}

	return r.mapFromEnt(subEnt), nil
}

// List implements subscription.Repository.
func (r *repositoryImpl) List(ctx context.Context, filter *subscription.SubscriptionFilter) ([]*subscription.Subscription, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Subscription.
		Query()

	if filter.IncludePlan {
		q = q.WithPlan()
	}

	if filter.IncludeTenant {
		q = q.WithTenant()
	}

	subs, err := q.All(ctx)
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
	var mappedPlan *plan.Plan
	if s.Edges.Plan != nil {
		mappedPlan = &plan.Plan{
			ID:          s.Edges.Plan.ID,
			Name:        *s.Edges.Plan.Name,
			Description: s.Edges.Plan.Description,
			MaxOrders:   s.Edges.Plan.MaxOrders,
			MaxUsers:    s.Edges.Plan.MaxUsers,
			Price:       s.Edges.Plan.Price,
			Duration:    s.Edges.Plan.DurationDays,
			CreatedAt:   s.Edges.Plan.CreatedAt,
			UpdatedAt:   s.Edges.Plan.UpdatedAt,
		}
	}

	var mappedTenant *tenant.Tenant
	if s.Edges.Tenant != nil {
		mappedTenant = &tenant.Tenant{
			ID:        s.Edges.Tenant.ID,
			Name:      *s.Edges.Tenant.Name,
			Phone:     *s.Edges.Tenant.Phone,
			Email:     *s.Edges.Tenant.Email,
			Address:   *s.Edges.Tenant.Address,
			CreatedAt: s.Edges.Tenant.CreatedAt,
			UpdatedAt: s.Edges.Tenant.UpdatedAt,
		}
	}

	return &subscription.Subscription{
		ID:        s.ID,
		PlanID:    s.PlanID,
		Plan:      mappedPlan,
		TenantID:  s.TenantID,
		Tenant:    mappedTenant,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		Status:    subscription.SubscriptionStatus(s.Status),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func (r *repositoryImpl) setActivePlan(ctx context.Context, tenantID uuid.UUID, planID uuid.UUID, endDate time.Time) error {
	cacheKey := redisutils.ActivePlan(tenantID)
	expiration := time.Until(endDate)
	if endDate.IsZero() {
		expiration = 0
	}

	return r.redisClient.Set(ctx, cacheKey, planID.String(), expiration).Err()
}

func NewRepository(client *db.Client, redisClient *db.RedisClient) subscription.Repository {
	return &repositoryImpl{
		client:      client,
		redisClient: redisClient,
	}
}
