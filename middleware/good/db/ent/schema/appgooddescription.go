//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppGoodDescription holds the schema definition for the AppGoodDescription entity.
type AppGoodDescription struct {
	ent.Schema
}

func (AppGoodDescription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

func (AppGoodDescription) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("description").
			Optional().
			Default(""),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the AppGoodDescription.
func (AppGoodDescription) Edges() []ent.Edge {
	return nil
}

func (AppGoodDescription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
