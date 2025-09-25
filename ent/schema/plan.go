package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.String("name").
			NotEmpty().
			Nillable().
			Unique(),

		field.Int("max_orders").
			Default(0).
			Nillable().
			Comment("0 means unlimited"),

		field.Int("max_users").
			Default(1).
			Nillable().
			Comment("0 means unlimited"),

		field.Float("price").
			Default(0.0).
			Nillable(),

		field.Int("duration_days").
			Default(30).
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("permissions", Permission.Type).
			Ref("plans"),

		edge.To("subscriptions", Subscription.Type),
	}
}
