package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Services holds the schema definition for the Services entity.
type Services struct {
	ent.Schema
}

// Fields of the Services.
func (Services) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),

		field.UUID("tenant_id", uuid.UUID{}).
			Optional().
			Nillable().
			Comment("Needed if user is not associated with a tenant"),

		field.UUID("category_id", uuid.UUID{}).
			Optional().
			Nillable(),

		field.String("name").
			NotEmpty().
			Nillable(),

		field.String("description").
			Optional().
			Nillable(),

		field.Float("base_price").
			Default(0.0).
			Nillable(),

		field.UUID("unit_id", uuid.UUID{}).
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

// Edges of the Services.
func (Services) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("services").
			Field("tenant_id").
			Unique(),

		edge.From("category", Category.Type).
			Ref("services").
			Field("category_id").
			Unique(),

		edge.From("unit", Unit.Type).
			Ref("services").
			Field("unit_id").
			Unique(),
	}
}
