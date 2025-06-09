package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FractionWithdrawalRule holds the schema definition for the FractionWithdrawalRule entity.
type FractionWithdrawalRule struct {
	ent.Schema
}

func (FractionWithdrawalRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FractionWithdrawalRule.
func (FractionWithdrawalRule) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("pool_coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("withdraw_interval").Optional().Default(0),
		field.
			Other("least_withdrawal_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("payout_threshold", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("withdraw_fee", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the FractionWithdrawalRule.
func (FractionWithdrawalRule) Edges() []ent.Edge {
	return nil
}
