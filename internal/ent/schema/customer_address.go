package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type CustomerAddress struct {
	ent.Schema
}

func (CustomerAddress) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("recipient_name").
			NotEmpty(),
		field.String("phone").
			Optional().
			Nillable(),
		field.Text("address").
			NotEmpty(),
		field.Float("latitude").
			Optional().
			Nillable(),
		field.Float("longitude").
			Optional().
			Nillable(),
		field.String("postal_code").
			Optional().
			Nillable(),
		field.Bool("is_default").
			Default(false),
	}
}

func (CustomerAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("addresses").
			Unique().
			Required(),
		edge.From("province", Province.Type).
			Ref("customer_addresses").
			Required(),
		edge.From("regency", Regency.Type).
			Ref("customer_addresses").
			Required(),
		edge.From("district", District.Type).
			Ref("customer_addresses").
			Required(),
		edge.To("pickup_orders", Order.Type),
		edge.To("delivery_orders", Order.Type),
	}
}
