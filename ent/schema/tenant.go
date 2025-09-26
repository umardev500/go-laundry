package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Tenant struct {
	ent.Schema
}

func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.String("name").
			NotEmpty().
			Nillable(),

		field.String("phone").
			NotEmpty().
			Nillable(),

		field.String("email").
			NotEmpty().
			Nillable(),

		field.String("address").
			NotEmpty().
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("customers", Customer.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("roles", Role.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("subscriptions", Subscription.Type),

		edge.To("tenant_usage", TenantUsage.Type),

		edge.To("payments", Payment.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}
