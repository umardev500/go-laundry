package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Province struct {
	ent.Schema
}

func (Province) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.String("name").
			NotEmpty(),
	}
}

func (Province) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("regencies", Regency.Type),
	}
}
