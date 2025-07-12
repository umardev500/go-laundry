package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("full_name").
			NotEmpty(),

		field.String("avatar_url").
			Optional().
			Nillable(),
	}
}

func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("customer").
			Unique().
			Required(),
		edge.To("spending", CustomerSpending.Type),
		edge.To("addresses", CustomerAddress.Type),
		edge.To("level_assignments", CustomerLevelAssignment.Type),
		edge.To("ratings", CustomerRating.Type),
		edge.To("orders", Order.Type),
		edge.To("payments", Payment.Type),
		edge.To("topups", Topup.Type),
	}
}
