package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CustomerLevelAssignment holds the schema definition for the CustomerLevelAssignment entity.
type CustomerLevelAssignment struct {
	ent.Schema
}

// Fields of the CustomerLevelAssignment.
func (CustomerLevelAssignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Time("assigned_at").
			Default(time.Now),
	}
}

// Edges of the CustomerLevelAssignment.
func (CustomerLevelAssignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("level_assignments").
			Unique().
			Required(),
		edge.From("tenant", Tenant.Type).
			Ref("customer_level_assignments").
			Unique().
			Required(),
		edge.From("level", CustomerLevel.Type).
			Ref("assignments").
			Unique().
			Required(),
	}
}
