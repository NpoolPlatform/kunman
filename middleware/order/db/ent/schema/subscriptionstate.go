package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// SubscriptionOrderState holds the schema definition for the SubscriptionOrderState entity.
type SubscriptionOrderState struct {
	ent.Schema
}

func (SubscriptionOrderState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the SubscriptionOrderState.
func (SubscriptionOrderState) Fields() []ent.Field {
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
		field.
			String("deal_event_id").
			Optional().
			Default(""),
	}
}

// Edges of the SubscriptionOrderState.
func (SubscriptionOrderState) Edges() []ent.Edge {
	return nil
}

func (SubscriptionOrderState) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
