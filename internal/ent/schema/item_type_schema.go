package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type ItemType struct {
	ent.Schema
}

func (ItemType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("name").
			Unique().
			Comment("Name of item type e.g. Dress, Shirt, Pants, etc"),
		field.Float("price").
			Default(0).
			Comment("Price of item type"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (ItemType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("unit", Unit.Type).
			Ref("item_types").
			Unique(),
		edge.To("order_items", OrderItem.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
