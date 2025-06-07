package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SubscriptionOrder holds the schema definition for the SubscriptionOrder entity.
type SubscriptionOrder struct {
	ent.Schema
}

func (SubscriptionOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the SubscriptionOrder.
func (SubscriptionOrder) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("good_value_usd", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("payment_amount_usd", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("discount_amount_usd", decimal.Decimal{}).
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
			Uint32("duration_seconds").
			Optional().
			Default(0),
	}
}

// Edges of the SubscriptionOrder.
func (SubscriptionOrder) Edges() []ent.Edge {
	return nil
}

func (SubscriptionOrder) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
