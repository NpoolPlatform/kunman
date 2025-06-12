package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Commission holds the schema definition for the Commission entity.
type Commission struct {
	ent.Schema
}

func (Commission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Commission.
func (Commission) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("amount_or_percent", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("start_at").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			String("settle_type").
			Optional().
			Default(types.SettleType_DefaultSettleType.String()),
		field.
			String("settle_mode").
			Optional().
			Default(types.SettleMode_DefaultSettleMode.String()),
		field.
			String("settle_interval").
			Optional().
			Default(types.SettleInterval_DefaultSettleInterval.String()),
		field.
			String("settle_amount_type").
			Optional().
			Default(types.SettleAmountType_SettleByPercent.String()),
		field.
			Other("threshold", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("order_limit").
			Optional().
			Default(0),
	}
}

// Edges of the Commission.
func (Commission) Edges() []ent.Edge {
	return nil
}

// Indexes of the Commission.
func (Commission) Indexes() []ent.Index {
	return nil
}
