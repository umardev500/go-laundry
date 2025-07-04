package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Plan struct {
	ent.Schema
}

func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("name").
			NotEmpty().
			Unique(),

		field.String("description").
			Optional().
			Nillable(),

		field.Float("price").
			Positive().
			Default(0),

		field.Int("max_branch").
			Optional().
			Nillable().
			Positive().
			Comment("Maximum number of branches.\nIf null, there is no limit."),

		field.Int("max_order").
			Optional().
			Nillable().
			Positive().
			Comment("Maximum number of orders.\nIf null, there is no limit."),

		field.Int("max_users").
			Optional().
			Nillable().
			Positive().
			Comment("Maximum number of admin users.\nIf null, there is no limit."),

		field.Int("max_customers").
			Optional().
			Nillable().
			Positive().
			Comment("Maximum number of customers.\nIf null, there is no limit."),

		field.Enum("billing_cycle").
			Values("monthly", "yearly").
			Default("monthly"),

		field.Int("duration_days").
			Min(1).
			Positive().
			Comment("Length of the billing cycle in days (e.g., 30 for monthly)."),

		field.Bool("enabled").
			Default(true),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("features", Feature.Type).
			Ref("plans"),
	}
}
