package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// GoodMalfunction holds the schema definition for the GoodMalfunction entity.
type GoodMalfunction struct {
	ent.Schema
}

func (GoodMalfunction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the GoodMalfunction.
func (GoodMalfunction) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Text("message").
			Optional().
			Default(""),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("duration_seconds").
			Optional().
			Default(0),
		field.
			Uint32("compensate_seconds").
			Optional().
			Default(0),
	}
}

// Edges of the GoodMalfunction.
func (GoodMalfunction) Edges() []ent.Edge {
	return nil
}

func (GoodMalfunction) Indexes() []ent.Index {
	return nil
}
