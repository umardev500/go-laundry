package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Merchant struct {
	ent.Schema
}

func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.String("name").
			NotEmpty().
			Comment("Merchant name e.g. Alpha Laundry, Beta Laundry, etc"),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("phone").
			Unique().
			NotEmpty(),
		field.String("address").
			NotEmpty(),
		field.Enum("status").
			Values("pending", "active", "inactive").
			Default("pending"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("guest_customers", GuestCustomers.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("item_types", ItemType.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("units", Unit.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("orders", Orders.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("roles", Role.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("merchant_features", MerchantFeature.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("subscriptions", Subscription.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
