package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("user_id", uuid.UUID{}).
			Immutable().
			Nillable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Immutable().
			Optional().
			Nillable(),

		field.UUID("reference_id", uuid.UUID{}).
			Immutable().
			Nillable().
			Optional(),

		field.Enum("reference_type").
			Values("subscription").
			Immutable().
			Nillable(),

		field.UUID("payment_method_id", uuid.UUID{}).
			Nillable(),

		field.String("proof_url").
			Optional().
			Nillable(),

		field.UUID("admin_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Admin who processed the payment"),

		// Details
		field.Float("amount").
			Default(0.0).
			Nillable(),

		field.Enum("currency").
			Values("IDR").
			Default("IDR").
			Nillable(),

		field.Enum("status").
			Values("pending", "completed", "failed", "cancelled").
			Default("pending").
			Nillable(),

		field.Time("paid_at").
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("payments").
			Field("user_id").
			Required().
			Immutable().
			Unique(),

		edge.From("tenant", Tenant.Type).
			Ref("payments").
			Field("tenant_id").
			Immutable().
			Unique(),

		edge.From("subscription", Subscription.Type).
			Ref("payments").
			Field("reference_id").
			Unique().
			Immutable().
			Comment("only used if reference_type is subscription"),

		edge.From("payment_method", PaymentMethod.Type).
			Ref("payments").
			Field("payment_method_id").
			Required().
			Unique(),

		edge.To("audit_logs", AuditLog.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
