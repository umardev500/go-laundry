package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// District holds the schema definition for the District entity.
type District struct {
	ent.Schema
}

// Fields of the District.
func (District) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(2).
			NotEmpty().
			Unique(),
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the District.
func (District) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("regency", Regency.Type).
			Ref("districts").
			Unique(),

		edge.To("villages", Village.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),

		edge.To("addresses", Addresses.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
