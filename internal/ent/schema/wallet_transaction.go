package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// WalletTransaction holds the schema definition for the WalletTransaction entity.
type WalletTransaction struct {
	ent.Schema
}

// Fields of the WalletTransaction.
func (WalletTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Float("amount"),

		field.Float("balance_after"),

		field.Enum("source_type").
			Values("topup", "subscription_payment"),

		field.UUID("source_id", uuid.UUID{}),

		field.String("notes").
			Optional().
			Nillable(),
	}
}

// Edges of the WalletTransaction.
func (WalletTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("wallet", Wallet.Type).
			Ref("transactions").
			Unique().
			Required(),
	}
}

// Indexes of the WalletTransaction.
func (WalletTransaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("wallet"),
	}
}
