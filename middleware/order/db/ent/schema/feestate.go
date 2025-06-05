package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// FeeOrderState holds the schema definition for the FeeOrderState entity.
type FeeOrderState struct {
	ent.Schema
}

func (FeeOrderState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FeeOrderState.
func (FeeOrderState) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("payment_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("paid_at").
			Optional().
			Default(0),
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
			String("payment_state").
			Optional().
			Default(types.PaymentState_PaymentStateWait.String()),
		field.
			String("cancel_state").
			Optional().
			Default(types.OrderState_DefaultOrderState.String()),
		field.
			Uint32("canceled_at").
			Optional().
			Default(0),
	}
}

// Edges of the FeeOrderState.
func (FeeOrderState) Edges() []ent.Edge {
	return nil
}

func (FeeOrderState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
