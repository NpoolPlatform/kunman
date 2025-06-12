package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppConfig holds the schema definition for the AppConfig entity.
type AppConfig struct {
	ent.Schema
}

func (AppConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppConfig.
func (AppConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("settle_mode").
			Optional().
			Default(types.SettleMode_DefaultSettleMode.String()),
		field.
			String("settle_amount_type").
			Optional().
			Default(types.SettleAmountType_SettleByPercent.String()),
		field.
			String("settle_interval").
			Optional().
			Default(types.SettleInterval_DefaultSettleInterval.String()),
		field.
			String("commission_type").
			Optional().
			Default(types.CommissionType_DefaultCommissionType.String()),
		field.
			Bool("settle_benefit").
			Optional().
			Default(false),
		field.
			Uint32("max_level").
			Optional().
			Default(1),
		field.
			Uint32("start_at").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
	}
}

// Edges of the AppConfig.
func (AppConfig) Edges() []ent.Edge {
	return nil
}

// Indexes of the AppConfig.
func (AppConfig) Indexes() []ent.Index {
	return nil
}
