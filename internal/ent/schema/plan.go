package schema

import (
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
			NotEmpty(),
		field.String("description").
			Optional().
			Nillable(), // allows NULL
	}
}

func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("variants", PlanVariant.Type),
	}
}
