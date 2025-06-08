package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppRole holds the schema definition for the AppRole entity.
type AppRole struct {
	ent.Schema
}

func (AppRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppRole.
func (AppRole) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("created_by", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("role").
			Optional().
			Default(""),
		field.
			String("description").
			Optional().
			Default(""),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("default").
			Optional().
			Default(false),
		field.
			Bool("genesis").
			Optional().
			Default(false),
	}
}

// Edges of the AppRole.
func (AppRole) Edges() []ent.Edge {
	return nil
}
