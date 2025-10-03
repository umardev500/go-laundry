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

		field.Enum("status").
			Values("active", "suspended", "deleted").
			Default("suspended").
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

		edge.To("audit_logs", AuditLog.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),

		edge.To("performed_audit_logs", AuditLog.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),

		edge.To("tenant_users", TenantUser.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),

		edge.To("platform_users", PlatformUser.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),
	}
}
