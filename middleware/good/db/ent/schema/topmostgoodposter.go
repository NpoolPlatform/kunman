//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// TopMostGoodPoster holds the schema definition for the TopMostGoodPoster entity.
type TopMostGoodPoster struct {
	ent.Schema
}

func (TopMostGoodPoster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

func (TopMostGoodPoster) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("top_most_good_id", uuid.UUID{}).
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

// Edges of the TopMostGoodPoster.
func (TopMostGoodPoster) Edges() []ent.Edge {
	return nil
}

func (TopMostGoodPoster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("top_most_good_id"),
	}
}
