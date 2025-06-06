package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Capacity holds the schema definition for the Capacity entity.
type Capacity struct {
	ent.Schema
}

func (Capacity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Capacity.
func (Capacity) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("capacity_key").
			Optional().
			Default(""),
		field.
			String("value").
			Optional().
			Default(""),
	}
}

// Edges of the Capacity.
func (Capacity) Edges() []ent.Edge {
	return nil
}
