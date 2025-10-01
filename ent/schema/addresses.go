package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Addresses holds the schema definition for the Addresses entity.
type Addresses struct {
	ent.Schema
}

// Fields of the Addresses.
func (Addresses) Fields() []ent.Field {
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

		field.String("zip_code").
			NotEmpty().
			Nillable(),

		field.Bool("is_default").
			Default(false),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Addresses.
func (Addresses) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("addresses").
			Field("user_id").
			Immutable().
			Unique().
			Required(),

		edge.From("tenant", Tenant.Type).
			Ref("addresses").
			Field("tenant_id").
			Immutable().
			Unique(),

		edge.From("province", Province.Type).
			Ref("addresses").
			Required().
			Unique(),

		edge.From("regency", Regency.Type).
			Ref("addresses").
			Required().
			Unique(),

		edge.From("district", District.Type).
			Ref("addresses").
			Required().
			Unique(),

		edge.From("village", Village.Type).
			Ref("addresses").
			Unique(),
	}
}
