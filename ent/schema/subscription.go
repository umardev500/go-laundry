package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("plan_id", uuid.UUID{}).
			Nillable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Nillable(),

		field.Time("start_date").
			Optional().
			Nillable(),

		field.Time("end_date").
			Optional().
			Nillable(),

		field.Enum("status").
			Values("active", "inactive", "pending", "cancelled", "suspended").
			Default("pending"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).
			Ref("subscriptions").
			Field("plan_id").
			Required().
			Unique(),

		edge.From("tenant", Tenant.Type).
			Ref("subscriptions").
			Field("tenant_id").
			Required().
			Unique(),

		edge.To("tenant_usage", TenantUsage.Type),

		edge.To("payments", Payment.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("audit_logs", AuditLog.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
