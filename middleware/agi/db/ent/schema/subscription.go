package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("next_extend_at").
			Optional().
			Default(0),
		field.
			Uint32("permanent_quota").
			Optional().
			Default(0),
		field.
			Uint32("consumed_quota").
			Optional().
			Default(0),
		field.
			Bool("pay_with_coin_balance").
			Optional().
			Default(false),
		field.
			String("subscription_id").
			Optional().
			Default(""),
		field.
			String("fiat_payment_channel").
			Optional().
			Default(""),
		field.
			Uint32("last_payment_at").
			Optional().
			Default(0),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return nil
}
