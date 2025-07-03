package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Orders struct {
	ent.Schema
}

func (Orders) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Enum("status").
			Values("pending", "ready_for_pickup", "on_the_way", "in_progress", "delivery", "completed", "cancelled").
			Default("pending"),
		field.Float("total_price").
			Default(0),
		field.String("pickup_address").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Orders) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("orders").
			Unique().
			Required(),
		edge.From("guest_customer", GuestCustomers.Type).
			Ref("orders").
			Unique(),
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
		edge.To("order_items", OrderItem.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
