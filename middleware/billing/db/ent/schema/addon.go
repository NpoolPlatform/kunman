package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Addon holds the schema definition for the Addon entity.
type Addon struct {
	ent.Schema
}

func (Addon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Addon.
func (Addon) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Other("usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("credit").
			Optional().
			Default(0),
		field.
			Uint32("sort_order").
			Optional().
			Default(0),
		field.
			Bool("enabled").
			Optional().
			Default(false),
		field.
			String("description").
			Optional().
			Default(""),
	}
}

// Edges of the Addon.
func (Addon) Edges() []ent.Edge {
	return nil
}
