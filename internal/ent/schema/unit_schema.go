package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Unit struct {
	ent.Schema
}

func (Unit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("name").
			Unique().
			NotEmpty().
			Comment("Unit name e.g. kilogram, meter, liter, etc"),
		field.String("description").
			Optional().
			Nillable().
			Comment("Description for unit e.g. kilogram is weight, meter is distance, liter is volume, etc"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Unit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("units").
			Unique().
			Required(),
		edge.To("item_types", ItemType.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
