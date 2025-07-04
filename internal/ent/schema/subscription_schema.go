package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Subscription struct {
	ent.Schema
}

func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Enum("status").
			Values("pending", "trial", "cancelled", "active", "inactive").
			Default("pending").
			Comment("Lifecycle status of the subscription."),

		field.Time("start_date").
			Default(time.Now).
			Comment("When the subscription starts."),

		field.Time("end_date").
			Optional().
			Nillable().
			Comment("When the subscription ends. Optional if ongoing."),

		field.Bool("auto_renew").
			Default(true).
			Comment("Indicates if the subscription will auto-renew."),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).
			Ref("subscriptions").
			Unique().
			Required(),
	}
}
