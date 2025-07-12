package schema

import (
	"entgo.io/ent"
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

		field.String("logo_url").
			Optional().
			Nillable(),

		field.Enum("status").
			Values("pending", "active", "suspended", "closed").
			Default("pending"),

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
	}
}

func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tenants").
			Unique().
			Required().
			Comment("The user that owns the tenant"),
		edge.From("plant_variant", PlanVariant.Type).
			Ref("tenants").
			Unique().
			Comment("The plan variant that owns the tenant"),
		edge.From("province", Province.Type).
			Ref("tenants").
			Required(),
		edge.From("regency", Regency.Type).
			Ref("tenants").
			Required(),
		edge.From("district", District.Type).
			Ref("tenants").
			Required(),
		edge.To("branches", Branch.Type),
		edge.To("tenant_users", TenantUser.Type),
		edge.To("roles", Role.Type),
		edge.To("customer_spending", CustomerSpending.Type),
		edge.To("customer_levels", CustomerLevel.Type),
		edge.To("customer_level_assignments", CustomerLevelAssignment.Type),
		edge.To("customer_ratings", CustomerRating.Type),
		edge.To("orders", Order.Type),
		edge.To("payments", Payment.Type),
		edge.To("payment_methods", PaymentMethod.Type),
		edge.To("subscriptions", Subscription.Type),
		edge.To("topups", Topup.Type),
		edge.To("wallet", Wallet.Type),
	}
}
