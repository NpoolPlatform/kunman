package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppSimulatePowerRental holds the schema definition for the AppSimulatePowerRental entity.
type AppSimulatePowerRental struct {
	ent.Schema
}

func (AppSimulatePowerRental) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

func (AppSimulatePowerRental) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional(),
		field.
			Other("order_units", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("order_duration_seconds").
			Optional().
			Default(0),
	}
}

// Edges of the AppSimulatePowerRental.
func (AppSimulatePowerRental) Edges() []ent.Edge {
	return nil
}
