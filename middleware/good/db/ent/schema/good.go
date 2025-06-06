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

// Good holds the schema definition for the Good entity.
type Good struct {
	ent.Schema
}

func (Good) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Good.
func (Good) Fields() []ent.Field {
	const benefitHours = 24
	return []ent.Field{
		field.
			UUID("device_info_id", uuid.UUID{}),
		field.
			UUID("coin_type_id", uuid.UUID{}),
		field.
			UUID("inherit_from_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("vendor_location_id", uuid.UUID{}),
		field.
			Other("unit_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("benefit_type").
			Optional().
			Default(types.BenefitType_DefaultBenefitType.String()),
		field.
			String("good_type").
			Optional().
			Default(types.GoodType_DefaultGoodType.String()),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			String("unit").
			Optional().
			Default(""),
		field.
			String("quantity_unit").
			Optional().
			Default(""),
		field.
			Int32("unit_amount").
			Optional().
			Default(0),
		field.
			Other("quantity_unit_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("delivery_at").
			Optional().
			Default(0),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			String("start_mode").
			Optional().
			Default(types.GoodStartMode_GoodStartModeNextDay.String()),
		field.
			Bool("test_only").
			Optional().
			Default(false),
		field.
			Uint32("benefit_interval_hours").
			Optional().
			Default(benefitHours),
		field.
			Other("unit_lock_deposit", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("unit_type").
			Optional().
			Default(types.GoodUnitType_GoodUnitByDurationAndQuantity.String()),
		field.
			String("quantity_calculate_type").
			Optional().
			Default(types.GoodUnitCalculateType_GoodUnitCalculateBySelf.String()),
		field.
			String("duration_type").
			Optional().
			Default(types.GoodDurationType_GoodDurationByDay.String()),
		field.
			String("duration_calculate_type").
			Optional().
			Default(types.GoodUnitCalculateType_GoodUnitCalculateBySelf.String()),
		field.
			String("settlement_type").
			Optional().
			Default(types.GoodSettlementType_GoodSettledByPaymentAmount.String()),
	}
}

// Edges of the Good.
func (Good) Edges() []ent.Edge {
	return nil
}

func (Good) Indexes() []ent.Index {
	return nil
}
