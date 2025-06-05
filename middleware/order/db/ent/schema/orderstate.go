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

// OrderState holds the schema definition for the OrderState entity.
type OrderState struct {
	ent.Schema
}

func (OrderState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OrderState.
func (OrderState) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			String("order_state").
			Optional().
			Default(types.OrderState_OrderStateCreated.String()),
		field.
			String("cancel_state").
			Optional().
			Default(types.OrderState_DefaultOrderState.String()),
		field.
			String("start_mode").
			Optional().
			Default(types.OrderStartMode_OrderStartConfirmed.String()),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			Uint32("paid_at").
			Optional().
			Default(0),
		field.
			Uint32("last_benefit_at").
			Optional().
			Default(0),
		field.
			String("benefit_state").
			Optional().
			Default(types.BenefitState_BenefitWait.String()),
		field.
			Bool("user_set_paid").
			Optional().
			Default(false),
		field.
			Bool("user_set_canceled").
			Optional().
			Default(false),
		field.
			Bool("admin_set_canceled").
			Optional().
			Default(false),
		field.
			String("payment_transaction_id").
			Optional().
			Default(""),
		field.
			Other("payment_finish_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("payment_state").
			Optional().
			Default(types.PaymentState_PaymentStateWait.String()),
		field.
			Uint32("outofgas_hours").
			Optional().
			Default(0),
		field.
			Uint32("compensate_hours").
			Optional().
			Default(0),
		field.
			String("renew_state").
			Optional().
			Default(types.OrderRenewState_OrderRenewWait.String()),
		field.
			Uint32("renew_notify_at").
			Optional().
			Default(0),
	}
}

// Edges of the OrderState.
func (OrderState) Edges() []ent.Edge {
	return nil
}

func (OrderState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
