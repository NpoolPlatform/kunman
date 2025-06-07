package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// GoodReward holds the schema definition for the GoodReward entity.
type GoodReward struct {
	ent.Schema
}

func (GoodReward) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the GoodReward.
func (GoodReward) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("reward_state").
			Optional().
			Default(types.BenefitState_BenefitWait.String()),
		field.
			Uint32("last_reward_at").
			Optional().
			Default(0),
	}
}

// Edges of the GoodReward.
func (GoodReward) Edges() []ent.Edge {
	return nil
}

func (GoodReward) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id"),
	}
}
