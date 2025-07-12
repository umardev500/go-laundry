package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Float("total_amount").
			Default(0.0),

		field.Float("total_discount_amount").
			Default(0.0),

		field.Float("final_amount").
			Positive(),

		field.Time("scheduled_pickup_at").
			Optional().
			Nillable(),

		field.Time("scheduled_delivery_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("orders").
			Unique().
			Required(),

		edge.From("tenant", Tenant.Type).
			Ref("orders").
			Unique().
			Required(),

		edge.From("branch", Branch.Type).
			Ref("orders").
			Unique().
			Required(),

		edge.From("pickup_address", CustomerAddress.Type).
			Ref("pickup_orders").
			Unique().
			Required(),

		edge.From("delivery_address", CustomerAddress.Type).
			Ref("delivery_orders").
			Unique().
			Required(),

		edge.From("payment", Payment.Type).
			Ref("orders").
			Unique().
			Required(),

		edge.To("customer_ratings", CustomerRating.Type),
	}
}
