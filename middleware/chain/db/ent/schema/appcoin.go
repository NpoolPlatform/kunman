package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppCoin holds the schema definition for the AppCoin entity.
type AppCoin struct {
	ent.Schema
}

func (AppCoin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppCoin.
func (AppCoin) Fields() []ent.Field {
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
			String("name").
			Optional().
			Default(""),
		field.
			JSON("display_names", []string{}).
			Optional().
			Default([]string{}),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			Bool("for_pay").
			Optional().
			Default(false),
		field.
			Other("withdraw_auto_review_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("product_page").
			Optional().
			Default(""),
		field.
			Bool("disabled").
			Optional().
			Default(false),
		field.
			Other("daily_reward_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("display").
			Optional().
			Default(true),
		field.
			Uint32("display_index").
			Optional().
			Default(0),
		field.
			Other("max_amount_per_withdraw", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the AppCoin.
func (AppCoin) Edges() []ent.Edge {
	return nil
}
