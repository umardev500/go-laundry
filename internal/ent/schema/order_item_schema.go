package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Float("amount").
			Comment("Dynamic quantity value according to the unit e.g. 1 kilogram, 2 liter, etc"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Orders.Type).
			Ref("order_items").
			Unique(),
		edge.From("item_type", ItemType.Type).
			Ref("order_items").
			Unique(),
	}
}
