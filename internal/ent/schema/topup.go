package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Topup holds the schema definition for the Topup entity.
type Topup struct {
	ent.Schema
}

// Fields of the Topup.
func (Topup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Float("amount"),

		field.Enum("status").
			Values("pending", "approved", "rejected").
			Default("pending"),

		field.Time("approved_at").
			Optional().
			Nillable(),

		field.String("notes").
			Optional().
			Nillable(),
	}
}

// Edges of the Topup.
func (Topup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("topups").
			Unique(),

		edge.From("customer", Customer.Type).
			Ref("topups").
			Unique(),
	}
}
