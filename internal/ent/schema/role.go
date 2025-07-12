package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("name").
			NotEmpty(),

		field.Enum("scope").
			Values("platform", "tenant").
			Default("platform"),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("roles").
			Unique(),
		edge.From("users", User.Type).
			Ref("roles").
			Required(),
		edge.From("permissions", Permission.Type).
			Ref("roles"),
	}
}
