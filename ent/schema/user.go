package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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

		field.UUID("tenant_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Needed if user is not associated with a tenant"),

		field.String("email").
			Unique().
			NotEmpty(),

		field.String("password").
			Sensitive().
			NotEmpty(),

		field.String("reset_token").
			Optional().
			Nillable(),

		field.Time("reset_expires_at").
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("users").
			Field("tenant_id").
			Unique(),

		edge.From("role", Role.Type).
			Ref("users"),

		edge.To("customers", Customer.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("profile", Profile.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("payments", Payment.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("addresses", Addresses.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
