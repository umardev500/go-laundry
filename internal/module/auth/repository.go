package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/auth"
	"github.com/umardev500/go-laundry/internal/utils/redisutils"
)

type repositoryImpl struct {
	redisClient *db.RedisClient
}

// GetActivePlanID implements auth.Repository.
func (r *repositoryImpl) GetActivePlanID(ctx context.Context, tenantID uuid.UUID) (uuid.UUID, error) {
	planID, err := r.redisClient.Get(ctx, redisutils.ActivePlan(tenantID)).Result()
	if err == nil {
		return uuid.Parse(planID)
	}

	return uuid.Nil, err
}

func NewRepository(redisClient *db.RedisClient) auth.Repository {
	return &repositoryImpl{
		redisClient: redisClient,
	}
}
