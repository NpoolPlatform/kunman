package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PubsubMessage holds the schema definition for the PubsubMessage entity.
type PubsubMessage struct {
	ent.Schema
}

// Fields of the PubsubMessage.
func (PubsubMessage) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("message_id").
			Optional().
			Default(""),
		field.
			String("state").
			Optional().
			Default(""),
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
			Text("arguments").
			Optional().
			Default(""),
	}
}

// Edges of the PubsubMessage.
func (PubsubMessage) Edges() []ent.Edge {
	return nil
}

func (PubsubMessage) Indexes() []ent.Index {
	return []ent.Index{}
}
