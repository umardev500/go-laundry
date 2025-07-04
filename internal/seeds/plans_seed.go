package seeds

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/ent/plan"
)

func SeedPlans(ctx context.Context, featureIDs []uuid.UUID, tx *ent.Tx) error {
	plans := []domain.CreatePlanRequest{
		{
			Name:         "Basic",
			Price:        1000,
			Description:  "Basic plan",
			MaxBranch:    1,
			MaxOrder:     10,
			MaxUsers:     5,
			MaxCustomers: 10,
			BillingCycle: plan.BillingCycleMonthly,
			DurationDays: 30,
			Enabled:      true,
		},
	}

	for _, p := range plans {
		err := tx.Plan.
			Create().
			AddFeatureIDs(featureIDs...).
			SetName(p.Name).
			SetPrice(p.Price).
			SetDescription(p.Description).
			SetMaxBranch(p.MaxBranch).
			SetMaxDailyOrders(p.MaxOrder).
			SetMaxUsers(p.MaxUsers).
			SetMaxCustomers(p.MaxCustomers).
			SetBillingCycle(p.BillingCycle).
			SetDurationDays(p.DurationDays).
			SetEnabled(p.Enabled).
			OnConflictColumns(plan.FieldName).
			UpdateNewValues().
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}
