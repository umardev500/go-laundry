package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Branch struct {
	ent.Schema
}

func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.String("name").
			NotEmpty(),

		field.String("description").
			Optional().
			Nillable(),

		field.String("phone").
			Optional().
			Nillable(),

		field.String("email").
			Optional().
			Nillable(),

		field.String("address").
			Optional().
			Nillable(),

		field.String("postal_code").
			Optional().
			Nillable(),

		field.Float("latitude").
			SchemaType(map[string]string{
				"mysql":    "decimal(10,6)",
				"postgres": "decimal(10,6)",
			}).
			Optional().
			Nillable(),

		field.Float("longitude").
			SchemaType(map[string]string{
				"mysql":    "decimal(10,6)",
				"postgres": "decimal(10,6)",
			}).
			Optional().
			Nillable(),

		field.Enum("status").
			Values("active", "suspended", "closed").
			Default("active"),
	}
}

func (Branch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("branches").
			Unique().
			Required(),

		edge.From("province", Province.Type).
			Ref("branches").
			Required(),

		edge.From("regency", Regency.Type).
			Ref("branches").
			Required(),

		edge.From("district", District.Type).
			Ref("branches").
			Required(),

		edge.To("tenant_users", TenantUser.Type),

		edge.To("orders", Order.Type),
		edge.To("customer_ratings", CustomerRating.Type),
		edge.To("payments", Payment.Type),
	}
}
