package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type CustomerLevel struct {
	ent.Schema
}

func (CustomerLevel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional().
			Nillable(),
		field.Float("min_total_spent").
			Positive().
			Default(0.0),
		field.Enum("discount_type").
			Values("fixed", "percentage").
			Default("fixed"),
		field.Float("discount_value").
			Positive().
			Default(0.0),
		field.Bool("is_active").
			Default(true),
	}
}

func (CustomerLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("customer_levels").
			Unique().
			Required(),
		edge.To("assignments", CustomerLevelAssignment.Type),
	}
}
