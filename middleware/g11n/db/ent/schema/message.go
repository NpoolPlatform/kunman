package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("lang_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("message_id").
			Optional().
			Default(""),
		field.
			Text("message").
			Optional().
			Default(""),
		field.
			Uint32("get_index").
			Optional().
			Default(0),
		field.
			Bool("disabled").
			Optional().
			Default(false),
		field.
			String("short").
			Optional().
			Default(""),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
