package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Village holds the schema definition for the Village entity.
type Village struct {
	ent.Schema
}

// Fields of the Village.
func (Village) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(2).
			NotEmpty().
			Unique(),
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the Village.
func (Village) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("district", District.Type).
			Ref("villages").
			Unique(),

		edge.To("addresses", Addresses.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
