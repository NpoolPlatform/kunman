package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppSubscription holds the schema definition for the AppSubscription entity.
type AppSubscription struct {
	ent.Schema
}

func (AppSubscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the AppSubscription.
func (AppSubscription) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Other("usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("product_id").
			Optional().
			Default(""),
		field.
			String("plan_id").
			Optional().
			Default(""),
		field.
			Uint32("trial_units").
			Optional().
			Default(1),
		field.
			Other("trial_usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			UUID("price_fiat_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Other("fiat_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("trial_fiat_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the AppSubscription.
func (AppSubscription) Edges() []ent.Edge {
	return nil
}

func (AppSubscription) Indexes() []ent.Index {
	return nil
}
