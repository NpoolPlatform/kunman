package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the Order.
//nolint:funlen
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("user_id", uuid.UUID{}),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("app_good_id", uuid.UUID{}),
		field.
			UUID("payment_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("parent_order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("units_v1", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("good_value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("good_value_usd", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("payment_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("discount_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			UUID("promotion_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("duration").
			Optional().
			Default(0),
		field.
			String("order_type").
			Optional().
			Default(types.OrderType_Normal.String()),
		field.
			String("investment_type").
			Optional().
			Default(types.InvestmentType_FullPayment.String()),
		field.
			JSON("coupon_ids", []uuid.UUID{}).
			Optional().
			Default(func() []uuid.UUID {
				return []uuid.UUID{}
			}),
		field.
			String("payment_type").
			Optional().
			Default(types.PaymentType_PayWithBalanceOnly.String()),
		field.
			UUID("coin_type_id", uuid.UUID{}),
		field.
			UUID("payment_coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("transfer_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("balance_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("coin_usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("local_coin_usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("live_coin_usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("simulate").
			Optional().
			Default(false),
		field.
			String("create_method").
			Optional().
			Default(types.OrderCreateMethod_OrderCreatedByPurchase.String()),
		field.
			Bool("multi_payment_coins").
			Optional().
			Default(false),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id", "good_id", "app_good_id"),
	}
}
