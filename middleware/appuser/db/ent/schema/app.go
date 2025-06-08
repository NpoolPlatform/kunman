package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("created_by", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("description").
			Optional().
			Default(""),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}
