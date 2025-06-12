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

// AppGoodCommissionConfig holds the schema definition for the AppGoodCommissionConfig entity.
type AppGoodCommissionConfig struct {
	ent.Schema
}

func (AppGoodCommissionConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppGoodCommissionConfig.
func (AppGoodCommissionConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("level").
			Optional().
			Default(0),
		field.
			Other("threshold_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
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
			Uint32("invites").
			Optional().
			Default(0),
		field.
			String("settle_type").
			Optional().
			Default(types.SettleType_DefaultSettleType.String()),
		field.
			Bool("disabled").
			Optional().
			Default(false),
	}
}

// Edges of the AppGoodCommissionConfig.
func (AppGoodCommissionConfig) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppGoodCommissionConfig.
func (AppGoodCommissionConfig) Indexes() []ent.Index {
	return nil
}
