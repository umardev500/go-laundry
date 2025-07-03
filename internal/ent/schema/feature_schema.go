package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Feature struct {
	ent.Schema
}

func (Feature) Fields() []ent.Field {
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
		field.Bool("enabled").
			Default(true),
	}
}

func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("merchant_features", MerchantFeature.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("plans", Plan.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
