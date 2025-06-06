//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// TopMostPoster holds the schema definition for the TopMostPoster entity.
type TopMostPoster struct {
	ent.Schema
}

func (TopMostPoster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

func (TopMostPoster) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("top_most_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("poster").
			Optional().
			Default(""),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the TopMostPoster.
func (TopMostPoster) Edges() []ent.Edge {
	return nil
}

func (TopMostPoster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("top_most_id"),
	}
}
