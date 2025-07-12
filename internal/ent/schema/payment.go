package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Float("amount"),

		field.Float("fee_amount").
			Default(0.0),

		field.Enum("method").
			Values("cash", "gateway"),

		field.JSON("metadata", map[string]interface{}{}).
			Optional(),

		field.Enum("status").
			Values("pending", "paid", "failed").
			Default("pending"),

		field.Enum("source_type").
			Values("subscription", "topup", "service_order"),

		field.UUID("source_id", uuid.UUID{}),

		field.Time("paid_at").
			Optional().
			Nillable(),

		field.String("notes").
			Optional().
			Nillable(),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment_method", PaymentMethod.Type).
			Ref("payments").
			Unique().
			Required(),
		edge.From("tenant", Tenant.Type).
			Ref("payments").
			Unique(),
		edge.From("customer", Customer.Type).
			Ref("payments").
			Unique(),
		edge.From("branch", Branch.Type).
			Ref("payments").
			Unique(),
		edge.To("orders", Order.Type),
	}
}
