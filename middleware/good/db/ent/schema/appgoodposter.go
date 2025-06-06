//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppGoodPoster holds the schema definition for the AppGoodPoster entity.
type AppGoodPoster struct {
	ent.Schema
}

func (AppGoodPoster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

func (AppGoodPoster) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
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

// Edges of the AppGoodPoster.
func (AppGoodPoster) Edges() []ent.Edge {
	return nil
}

func (AppGoodPoster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
