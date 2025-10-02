package plan

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	permissionEntity "github.com/umardev500/go-laundry/ent/permission"
	planEntity "github.com/umardev500/go-laundry/ent/plan"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/permission"
	"github.com/umardev500/go-laundry/internal/domain/plan"
)

type repositoryImpl struct {
	client      *db.Client
	redisClient *db.RedisClient
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

	if err != nil {
		return fmt.Errorf("failed to replace permissions: %w", err)
	}

	// --- Redis cache update ---
	if r.redisClient != nil {
		cacheKey := fmt.Sprintf("plan:%s:permissions", planID)

		// Convert permissions to names
		names := make([]string, len(perms))
		for i, p := range perms {
			names[i] = *p.Name
		}

		// Use a transaction to replace the Redis Set atomically
		pipe := r.redisClient.TxPipeline()
		pipe.Del(ctx, cacheKey) // remove old permissions
		if len(names) > 0 {
			pipe.SAdd(ctx, cacheKey, names) // add new permissions
		}
		if _, err := pipe.Exec(ctx); err != nil {
			fmt.Printf("failed to update redis cache for replace permissions: %v\n", err)
		}
	}

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

	if err != nil {
		return fmt.Errorf("failed to remove permissions: %w", err)
	}

	// --- Redis cache update ---
	if r.redisClient != nil {
		cacheKey := fmt.Sprintf("plan:%s:permissions", planID)

		// Convert permissions to names
		names := make([]string, len(perms))
		for i, p := range perms {
			names[i] = *p.Name
		}

		// Remove from Redis Set
		if err := r.redisClient.SRem(ctx, cacheKey, names).Err(); err != nil {
			// Log error but don’t block main flow
			fmt.Printf("failed to remove permissions from redis: %v\n", err)
		}
	}

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
	if err != nil {
		return fmt.Errorf("failed to add permissions: %w", err)
	}

	// --- Redis caching step ---
	cacheKey := fmt.Sprintf("plan:%s:permissions", plantEnt.ID)

	// Convert permissions to names
	names := make([]string, len(perms))
	for i, p := range perms {
		names[i] = *p.Name
	}

	// Add to a Redis Set (avoids duplicates automatically)
	if err := r.redisClient.SAdd(ctx, cacheKey, names).Err(); err != nil {
		return fmt.Errorf("failed to cache permissions: %w", err)
	}

	return nil
}

// GetByID implements plan.Repository.
func (r *repositoryImpl) GetByID(ctx context.Context, id uuid.UUID, filter *plan.Filter) (*plan.Plan, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Plan.Query()

	if !filter.IncludeDeleted {
		q = q.Where(planEntity.DeletedAtIsNil())
	}

	if filter.IncludePermissions {
		q = q.WithPermissions()
	}

	entPlan, err := q.Where(planEntity.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.mapFromEnt(entPlan), nil
}

// List implements plan.Repository.
func (r *repositoryImpl) List(ctx context.Context, filter *plan.Filter) ([]*plan.Plan, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Plan.Query()

	if !filter.IncludeDeleted {
		q = q.Where(planEntity.DeletedAtIsNil())
	}

	if filter.IncludePermissions {
		q = q.WithPermissions()
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
	var mappedPermissions []*permission.Permission
	if e.Edges.Permissions != nil {
		for _, p := range e.Edges.Permissions {
			mappedPermissions = append(mappedPermissions, &permission.Permission{
				ID:   p.ID,
				Name: *p.Name,
			})
		}
	}

	return &plan.Plan{
		ID:          e.ID,
		Name:        *e.Name,
		Description: e.Description,
		MaxOrders:   e.MaxOrders,
		MaxUsers:    e.MaxUsers,
		Price:       e.Price,
		Duration:    e.DurationDays,
		Permissions: mappedPermissions,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func NewRepository(client *db.Client, redisClient *db.RedisClient) plan.Repository {
	return &repositoryImpl{
		client:      client,
		redisClient: redisClient,
	}
}
