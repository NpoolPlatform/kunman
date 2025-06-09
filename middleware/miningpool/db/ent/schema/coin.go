package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Coin holds the schema definition for the Coin entity.
type Coin struct {
	ent.Schema
}

func (Coin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Coin.
func (Coin) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("pool_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("coin_type").Optional().Default(""),
		field.
			Other("fee_ratio", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("fixed_revenue_able").Optional().Default(false),
		field.
			Other("least_transfer_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("benefit_interval_seconds").Optional().Default(0),
		field.
			String("remark").Optional().Default(""),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return nil
}

func (Coin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pool_id", "coin_type_id"),
	}
}
