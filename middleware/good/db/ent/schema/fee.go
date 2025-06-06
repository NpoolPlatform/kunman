package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Fee holds the schema definition for the Fee entity.
type Fee struct {
	ent.Schema
}

func (Fee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Fee.
func (Fee) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("settlement_type").
			Optional().
			Default(types.GoodSettlementType_GoodSettledByPaymentAmount.String()),
		// Unit value with amount or percent
		field.
			Other("unit_value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("duration_display_type").
			Optional().
			Default(types.GoodDurationType_GoodDurationByDay.String()),
	}
}

// Edges of the Fee.
func (Fee) Edges() []ent.Edge {
	return nil
}

func (Fee) Indexes() []ent.Index {
	return nil
}
