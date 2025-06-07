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

// PaymentFiat holds the schema definition for the PaymentFiat entity.
type PaymentFiat struct {
	ent.Schema
}

func (PaymentFiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the PaymentFiat.
func (PaymentFiat) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("payment_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("payment_channel").
			Optional().
			Default(types.FiatPaymentChannel_PaymentChannelStripe.String()),
		field.
			Other("amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("usd_currency", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the PaymentFiat.
func (PaymentFiat) Edges() []ent.Edge {
	return nil
}

func (PaymentFiat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("payment_id"),
	}
}
