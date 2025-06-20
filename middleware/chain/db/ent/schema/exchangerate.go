package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExchangeRate holds the schema definition for the ExchangeRate entity.
type ExchangeRate struct {
	ent.Schema
}

func (ExchangeRate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the ExchangeRate.
func (ExchangeRate) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			Other("market_value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("settle_value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("settle_percent").
			Optional().
			Default(100),
		field.
			JSON("settle_tips", []string{}).
			Optional().
			Default([]string{}),
		field.
			UUID("setter", uuid.UUID{}).
			Optional().
			Default(uuid.New),
	}
}

// Edges of the ExchangeRate.
func (ExchangeRate) Edges() []ent.Edge {
	return nil
}
