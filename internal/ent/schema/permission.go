package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Permission struct {
	ent.Schema
}

func (Permission) Fields() []ent.Field {
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
	}
}

func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feature", Feature.Type).
			Ref("permissions").
			Unique().
			Required(),
		edge.To("roles", Role.Type),
		edge.To("plant_variants", PlanVariant.Type),
	}
}
