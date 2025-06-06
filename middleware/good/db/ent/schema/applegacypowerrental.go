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

// AppLegacyPowerRental holds the schema definition for the AppLegacyPowerRental entity.
type AppLegacyPowerRental struct {
	ent.Schema
}

func (AppLegacyPowerRental) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

func (AppLegacyPowerRental) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("technique_fee_ratio", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the AppLegacyPowerRental.
func (AppLegacyPowerRental) Edges() []ent.Edge {
	return nil
}

func (AppLegacyPowerRental) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
