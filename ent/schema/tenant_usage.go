package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TenantUsage holds the schema definition for the TenantUsage entity.
type TenantUsage struct {
	ent.Schema
}

// Fields of the TenantUsage.
func (TenantUsage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Nillable(),

		field.UUID("subscription_id", uuid.UUID{}).
			Nillable(),

		field.Int("orders_count").
			Default(0).
			Comment("Number of orders this tenant has used in this period"),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the TenantUsage.
func (TenantUsage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("tenant_usage").
			Field("tenant_id").
			Required().
			Unique(),
		edge.From("subscription", Subscription.Type).
			Ref("tenant_usage").
			Field("subscription_id").
			Required().
			Unique(),
	}
}
