package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// SendAnnouncement holds the schema definition for the SendAnnouncement entity.
type SendAnnouncement struct {
	ent.Schema
}

func (SendAnnouncement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the SendAnnouncement.
func (SendAnnouncement) Fields() []ent.Field {
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
			UUID("announcement_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("channel").
			Optional().
			Default(basetypes.NotifChannel_DefaultChannel.String()),
	}
}

// Edges of the SendAnnouncement.
func (SendAnnouncement) Edges() []ent.Edge {
	return nil
}
