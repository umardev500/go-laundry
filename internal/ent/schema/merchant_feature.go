package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type MerchantFeature struct {
	ent.Schema
}

func (MerchantFeature) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Bool("enabled").
			Default(true),
	}
}

func (MerchantFeature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feature", Feature.Type).
			Ref("merchant_features").
			Unique().
			Required(),
		edge.From("merchant", Merchant.Type).
			Ref("merchant_features").
			Unique().
			Required(),
	}
}
