//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// PaymentBalanceLock holds the schema definition for the PaymentBalanceLock entity.
type PaymentBalanceLock struct {
	ent.Schema
}

func (PaymentBalanceLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the PaymentBalanceLock.
func (PaymentBalanceLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("payment_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("ledger_lock_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
	}
}

// Edges of the PaymentBalanceLock.
func (PaymentBalanceLock) Edges() []ent.Edge {
	return nil
}

func (PaymentBalanceLock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("payment_id"),
	}
}
