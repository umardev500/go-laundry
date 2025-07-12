package schema

import (
	"entgo.io/ent"
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
			Immutable().
			Unique(),

		field.Float("amount"),

		field.Enum("status").
			Values("pending", "active", "suspended", "canceled"),

		field.Time("start_date").
			Optional().
			Nillable(),

		field.Time("end_date").
			Optional().
			Nillable(),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("subscriptions").
			Unique().
			Required(),

		edge.From("plan_variant", PlanVariant.Type).
			Ref("subscriptions").
			Unique().
			Required(),
	}
}
