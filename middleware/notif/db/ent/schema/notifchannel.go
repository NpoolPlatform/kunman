//nolint:nolintlint,dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

// NotifChannel holds the schema definition for the NotifChannel entity.
type NotifChannel struct {
	ent.Schema
}

func (NotifChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the NotifChannel.
func (NotifChannel) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("event_type").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			String("channel").
			Optional().
			Default(basetypes.NotifChannel_DefaultChannel.String()),
	}
}

// Edges of the NotifChannel.
func (NotifChannel) Edges() []ent.Edge {
	return nil
}
