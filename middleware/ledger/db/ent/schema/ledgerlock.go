package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// LedgerLock holds the schema definition for the LedgerLock entity.
type LedgerLock struct {
	ent.Schema
}

func (LedgerLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the LedgerLock.
func (LedgerLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("ledger_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("statement_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			String("lock_state").
			Optional().
			Default(types.LedgerLockState_LedgerLockLocked.String()),
		// To support one lock for multi coins
		// In default it'll be set to ent id
		field.
			UUID("ex_lock_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
	}
}

// Edges of the LedgerLock.
func (LedgerLock) Edges() []ent.Edge {
	return nil
}

func (LedgerLock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ex_lock_id"),
	}
}
