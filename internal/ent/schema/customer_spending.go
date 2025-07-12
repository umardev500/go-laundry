package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type CustomerSpending struct {
	ent.Schema
}

func (CustomerSpending) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Float("total_spent").
			Default(0.0).
			SchemaType(map[string]string{
				"mysql":    "decimal(10,2)",
				"postgres": "decimal(10,2)",
			}),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (CustomerSpending) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("spending").
			Unique().
			Required(),

		edge.From("tenant", Tenant.Type).
			Ref("customer_spending").
			Unique().
			Required(),
	}
}

func (CustomerSpending) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("customer", "tenant").Unique(),
	}
}
