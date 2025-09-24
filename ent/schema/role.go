package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
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

		field.String("description").
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("roles").
			Field("tenant_id").
			Unique(),
		edge.From("permissions", Permission.Type).
			Ref("roles"),
		edge.To("users", User.Type),
	}
}
