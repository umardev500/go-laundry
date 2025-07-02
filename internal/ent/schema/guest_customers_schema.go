package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type GuestCustomers struct {
	ent.Schema
}

func (GuestCustomers) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("name").
			NotEmpty(),
		field.String("phone").
			NotEmpty(),
		field.String("address").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (GuestCustomers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Orders.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
