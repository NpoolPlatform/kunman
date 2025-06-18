package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

// Notif holds the schema definition for the Notif entity.
type Notif struct {
	ent.Schema
}

func (Notif) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Notif.
func (Notif) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("notified").
			Optional().
			Default(false),
		field.
			UUID("lang_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("event_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("event_type").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			Bool("use_template").
			Optional().
			Default(false),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Text("content").
			Optional().
			Default(""),
		field.
			String("channel").
			Optional().
			Default(basetypes.NotifChannel_DefaultChannel.String()),
		field.
			Text("extra").
			Optional().
			Default(""),
		field.
			String("type").
			Optional().
			Default(basetypes.NotifType_DefaultNotifType.String()),
	}
}

// Edges of the Notif.
func (Notif) Edges() []ent.Edge {
	return nil
}
