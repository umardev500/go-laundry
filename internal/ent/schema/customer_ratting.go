package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// CustomerRating holds the schema definition for the CustomerRating entity.
type CustomerRating struct {
	ent.Schema
}

// Fields of the CustomerRating.
func (CustomerRating) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Int("rating").
			Positive().
			Range(1, 5),

		field.String("comment").
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now),

		field.String("response").
			Optional().
			Nillable(),

		field.Time("responded_at").
			Optional().
			Nillable(),

		field.Bool("is_visible").
			Default(true),
	}
}

// Edges of the CustomerRating.
func (CustomerRating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("ratings").
			Unique().
			Required(),

		edge.From("tenant", Tenant.Type).
			Ref("customer_ratings").
			Unique().
			Required(),

		edge.From("branch", Branch.Type).
			Ref("customer_ratings").
			Unique().
			Required(),

		edge.From("order", Order.Type).
			Ref("customer_ratings").
			Unique().
			Required(),

		edge.From("user", User.Type).
			Ref("customer_ratings").
			Unique(),
	}
}

// Indexes of the CustomerRating.
func (CustomerRating) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("customer", "order").
			Unique(),
	}
}
