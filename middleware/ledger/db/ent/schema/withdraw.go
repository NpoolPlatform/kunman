package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
)

// Withdraw holds the schema definition for the Withdraw entity.
type Withdraw struct {
	ent.Schema
}

func (Withdraw) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Withdraw.
func (Withdraw) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("address").
			Optional().
			Default(""),
		field.
			UUID("platform_transaction_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("chain_transaction_id").
			Optional().
			Default(""),
		field.
			String("state").
			Optional().
			Default(basetypes.WithdrawState_Created.String()),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			UUID("review_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the Withdraw.
func (Withdraw) Edges() []ent.Edge {
	return nil
}
