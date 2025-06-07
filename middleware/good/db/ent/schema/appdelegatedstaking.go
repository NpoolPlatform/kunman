package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppDelegatedStaking holds the schema definition for the AppDelegatedStaking entity.
type AppDelegatedStaking struct {
	ent.Schema
}

func (AppDelegatedStaking) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

func (AppDelegatedStaking) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("service_start_at").
			Optional().
			Default(0),
		field.
			String("start_mode").
			Optional().
			Default(types.GoodStartMode_GoodStartModeNextDay.String()),
		field.
			Bool("enable_set_commission").
			Optional().
			Default(false),
	}
}

// Edges of the AppDelegatedStaking.
func (AppDelegatedStaking) Edges() []ent.Edge {
	return nil
}

func (AppDelegatedStaking) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
