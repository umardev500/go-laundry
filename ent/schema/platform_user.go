package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PlatformUser holds the schema definition for the PlatformUser entity.
type PlatformUser struct {
	ent.Schema
}

// Fields of the PlatformUser.
func (PlatformUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("user_id", uuid.UUID{}).
			Immutable().
			Nillable(),

		field.Enum("status").
			Values("active", "suspended", "deleted").
			Default("suspended").
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the PlatformUser.
func (PlatformUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("platform_users").
			Field("user_id").
			Immutable().
			Required().
			Unique(),

		edge.From("role", Role.Type).
			Ref("platform_users"),
	}
}
