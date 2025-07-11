package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Regency struct {
	ent.Schema
}

func (Regency) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.String("name").
			NotEmpty(),
	}
}

func (Regency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("province", Province.Type).
			Ref("regencies").
			Required(),
		edge.To("districts", District.Type),
	}
}
