//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinFiatCurrency holds the schema definition for the CoinFiatCurrency entity.
type CoinFiatCurrency struct {
	ent.Schema
}

func (CoinFiatCurrency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CoinFiatCurrency.
func (CoinFiatCurrency) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(basetypes.CurrencyFeedType_DefaultFeedType.String()),
		field.
			Other("market_value_low", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("market_value_high", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the CoinFiatCurrency.
func (CoinFiatCurrency) Edges() []ent.Edge {
	return nil
}
