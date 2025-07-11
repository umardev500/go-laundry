package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type PlanVariant struct {
	ent.Schema
}

func (PlanVariant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Enum("billing_cycle").
			Values("monthly", "quarterly", "yearly"),
		field.Float("price").
			SchemaType(map[string]string{
				"mysql":    "decimal(10,2)",
				"postgres": "decimal(10,2)",
			}).
			Positive(),
		field.Bool("is_active").
			Default(true),
	}
}

func (PlanVariant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", Plan.Type).
			Ref("variants").
			Unique(),
	}
}
