package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// TenantUser holds the schema definition for the TenantUser entity.
type TenantUser struct {
	ent.Schema
}

// Fields of the TenantUser.
func (TenantUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Immutable().
			Nillable(),

		field.UUID("user_id", uuid.UUID{}).
			Immutable().
			Nillable(),

		field.Enum("status").
			Values("active", "suspended", "deleted").
			Default("active").
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the TenantUser.
func (TenantUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("tenant_users").
			Field("tenant_id").
			Required().
			Immutable().
			Unique(),

		edge.From("user", User.Type).
			Ref("tenant_users").
			Field("user_id").
			Required().
			Immutable().
			Unique(),

		edge.From("role", Role.Type).
			Ref("tenant_users"),
	}
}

// Index
func (TenantUser) Indexes() []ent.Index {
	return []ent.Index{
		// 👇 Composite unique constraint: (tenant_id, user_id)
		index.Fields("tenant_id", "user_id").
			Unique(),
	}
}
