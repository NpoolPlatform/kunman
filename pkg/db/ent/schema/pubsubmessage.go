package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	msgpb "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// PubsubMessage holds the schema definition for the PubsubMessage entity.
type PubsubMessage struct {
	ent.Schema
}

func (PubsubMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the PubsubMessage.
func (PubsubMessage) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("message_id").
			Optional().
			Default(msgpb.MsgID_DefaultMsgID.String()),
		field.
			String("state").
			Optional().
			Default(msgpb.MsgState_DefaultMsgState.String()),
		field.
			UUID("resp_to_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("undo_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("arguments").
			Optional().
			Default(""),
	}
}

// Edges of the PubsubMessage.
func (PubsubMessage) Edges() []ent.Edge {
	return nil
}

func (PubsubMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("state", "resp_to_id"),
		index.Fields("state", "undo_id"),
	}
}
