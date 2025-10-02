package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuditLog holds the schema definition for the AuditLog entity.
type AuditLog struct {
	ent.Schema
}

// Fields of the AuditLog.
func (AuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.Enum("table_name").
			Values("payment", "subscription", "user", "tenant").
			Nillable().
			Comment("Table being changed"),

		field.UUID("record_id", uuid.UUID{}).
			Nillable().
			Comment("Primary key of the record"),

		field.Enum("action").
			Values("create", "update", "delete").
			Nillable(),

		field.UUID("modified_by", uuid.UUID{}).
			Nillable().
			Comment("User who made the change"),

		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.JSON("old_data", map[string]any{}).
			Optional(),

		field.JSON("new_data", map[string]any{}).
			Optional(),
	}
}

// Edges of the AuditLog.
func (AuditLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("performed_audit_logs").
			Field("modified_by").
			Unique().
			Required(),

		// Generic user edge without specific field
		edge.From("user_generic", User.Type).
			Ref("audit_logs").
			Unique(),

		edge.From("tenant", Tenant.Type).
			Ref("audit_logs").
			Unique(),

		edge.From("subscription", Subscription.Type).
			Ref("audit_logs").
			Unique(),

		edge.From("payment", Payment.Type).
			Ref("audit_logs").
			Unique(),
	}
}
