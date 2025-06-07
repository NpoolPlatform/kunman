package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppDefaultGood holds the schema definition for the AppDefaultGood entity.
type AppDefaultGood struct {
	ent.Schema
}

func (AppDefaultGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

func (AppDefaultGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the AppDefaultGood.
func (AppDefaultGood) Edges() []ent.Edge {
	return nil
}
