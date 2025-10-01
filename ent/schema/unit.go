package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Unit holds the schema definition for the Unit entity.
type Unit struct {
	ent.Schema
}

// Fields of the Unit.
func (Unit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Needed if user is not associated with a tenant"),

		field.String("name").
			NotEmpty().
			Nillable(),
	}
}

// Edges of the Unit.
func (Unit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("units").
			Field("tenant_id").
			Unique(),

		edge.To("services", Services.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
