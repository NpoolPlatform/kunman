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

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("good_type").
			Optional().
			Default(types.GoodType_Subscription.String()),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("duration_display_type").
			Optional().
			Default(types.GoodDurationType_GoodDurationByMonth.String()),
		field.
			Uint32("duration_units").
			Optional().
			Default(1),
		field.
			Uint32("duration_quota").
			Optional().
			Default(1080),
		field.
			Uint32("daily_bonus_quota").
			Optional().
			Default(0),
		field.
			Other("usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return nil
}

func (Subscription) Indexes() []ent.Index {
	return nil
}
