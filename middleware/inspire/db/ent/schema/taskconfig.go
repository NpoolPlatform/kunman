package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// TaskConfig holds the schema definition for the TaskConfig entity.
type TaskConfig struct {
	ent.Schema
}

func (TaskConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the TaskConfig.
func (TaskConfig) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("event_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("task_type").
			Optional().
			Default(types.TaskType_DefaultTaskType.String()),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("task_desc").
			Optional().
			Default(""),
		field.
			String("step_guide").
			Optional().
			Default(""),
		field.
			String("recommend_message").
			Optional().
			Default(""),
		field.
			Uint32("index").
			Optional().
			Default(0),
		field.
			UUID("last_task_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("max_reward_count").
			Optional().
			Default(0),
		field.
			Uint32("cooldown_second").
			Optional().
			Default(0),
		field.
			Bool("interval_reset").
			Optional().
			Default(false),
		field.
			Uint32("interval_reset_second").
			Optional().
			Default(0),
		field.
			Uint32("max_interval_reward_count").
			Optional().
			Default(0),
	}
}

// Edges of the TaskConfig.
func (TaskConfig) Edges() []ent.Edge {
	return nil
}
