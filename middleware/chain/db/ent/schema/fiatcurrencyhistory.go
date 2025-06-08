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

// FiatCurrencyHistory holds the schema definition for the FiatCurrencyHistory entity.
type FiatCurrencyHistory struct {
	ent.Schema
}

func (FiatCurrencyHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FiatCurrencyHistory.
func (FiatCurrencyHistory) Fields() []ent.Field {
	return []ent.Field{
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

// Edges of the FiatCurrencyHistory.
func (FiatCurrencyHistory) Edges() []ent.Edge {
	return nil
}
