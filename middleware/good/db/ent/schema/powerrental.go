package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PowerRental holds the schema definition for the PowerRental entity.
type PowerRental struct {
	ent.Schema
}

func (PowerRental) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the PowerRental.
func (PowerRental) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("device_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("vendor_location_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("unit_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("quantity_unit").
			Optional().
			Default(""),
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
			Other("unit_lock_deposit", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("duration_display_type").
			Optional().
			Default(types.GoodDurationType_GoodDurationByDay.String()),
		field.
			String("stock_mode").
			Optional().
			Default(types.GoodStockMode_GoodStockByUnique.String()),
	}
}

// Edges of the PowerRental.
func (PowerRental) Edges() []ent.Edge {
	return nil
}

func (PowerRental) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id"),
	}
}
