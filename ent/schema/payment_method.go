package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PaymentMethod holds the schema definition for the PaymentMethod entity.
type PaymentMethod struct {
	ent.Schema
}

// Fields of the PaymentMethod.
func (PaymentMethod) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Needed if user is not associated with a tenant"),

		field.UUID("payment_method_type_id", uuid.UUID{}).
			Nillable(),

		field.JSON("metadata", map[string]any{}).
			Optional(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the PaymentMethod.
func (PaymentMethod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payments", Payment.Type),

		edge.From("tenant", Tenant.Type).
			Ref("payment_methods").
			Field("tenant_id").
			Unique(),

		edge.From("payment_method_type", PaymentMethodType.Type).
			Ref("payment_methods").
			Field("payment_method_type_id").
			Required().
			Unique(),
	}
}
