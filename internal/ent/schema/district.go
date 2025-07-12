package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type District struct {
	ent.Schema
}

func (District) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique(),
		field.String("name").
			NotEmpty(),
	}
}

func (District) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("regency", Regency.Type).
			Ref("districts").
			Required(),
		edge.To("tenants", Tenant.Type),
		edge.To("branches", Branch.Type),
		edge.To("customer_addresses", CustomerAddress.Type),
	}
}
