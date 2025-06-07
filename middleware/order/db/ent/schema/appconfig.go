package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppConfig holds the schema definition for the AppConfig entity.
type AppConfig struct {
	ent.Schema
}

func (AppConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the AppConfig.
func (AppConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Bool("enable_simulate_order").
			Optional().
			Default(false),
		field.
			String("simulate_order_coupon_mode").
			Optional().
			Default(types.SimulateOrderCouponMode_WithoutCoupon.String()),
		field.
			Other("simulate_order_coupon_probability", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Other("simulate_order_cashable_profit_probability", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Uint32("max_unpaid_orders").
			Optional().
			Default(5), //nolint
		field.
			Uint32("max_typed_coupons_per_order").
			Optional().
			Default(1),
	}
}

// Edges of the AppConfig.
func (AppConfig) Edges() []ent.Edge {
	return nil
}
