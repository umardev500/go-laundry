package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type TenantUser struct {
	ent.Schema
}

func (TenantUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.String("name").
			NotEmpty(),

		field.String("avatar_url").
			Optional().
			Nillable(),

		field.Enum("status").
			Values("active", "inactive", "suspended").
			Default("active"),
	}
}

func (TenantUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tenant_users").
			Unique().
			Required(),

		edge.From("tenant", Tenant.Type).
			Ref("tenant_users").
			Unique().
			Required(),

		edge.From("branch", Branch.Type).
			Ref("tenant_users").
			Unique().
			Required(),
	}
}

func (TenantUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("user", "tenant", "branch").Unique(),
	}
}
