package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Regency holds the schema definition for the Regency entity.
type Regency struct {
	ent.Schema
}

// Fields of the Regency.
func (Regency) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(2).
			NotEmpty().
			Unique(),

		field.String("province_id").
			MaxLen(2).
			NotEmpty().
			Unique(),

		field.String("name").
			NotEmpty(),
	}
}

// Edges of the Regency.
func (Regency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("province", Province.Type).
			Ref("regencies").
			Field("province_id").
			Required().
			Unique(),

		edge.To("districts", District.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),

		edge.To("addresses", Addresses.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
