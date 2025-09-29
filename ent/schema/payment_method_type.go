package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PaymentMethodType holds the schema definition for the PaymentMethodType entity.
type PaymentMethodType struct {
	ent.Schema
}

// Fields of the PaymentMethodType.
func (PaymentMethodType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.String("name").
			NotEmpty().
			Nillable().
			Unique().
			Comment("Name of the payment method type e.g credit_card, bank_transfer"),

		field.String("display_name").
			NotEmpty().
			Nillable().
			Comment("Display name of the payment method type e.g Credit Card, Bank Transfer"),

		field.Enum("status").
			Values("active", "inactive").
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

// Edges of the PaymentMethodType.
func (PaymentMethodType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("payment_methods", PaymentMethod.Type),
	}
}
