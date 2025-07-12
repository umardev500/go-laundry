package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("phone").
			NotEmpty().
			Unique(),
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("password_hash").
			NotEmpty(),
		field.Enum("type").
			Values("platform", "tenant", "customer"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenants", Tenant.Type),
		edge.To("tenant_users", TenantUser.Type),
		edge.To("roles", Role.Type),
		edge.To("customer", Customer.Type),
		edge.To("customer_ratings", CustomerRating.Type),
	}
}
