package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppFiat holds the schema definition for the AppFiat entity.
type AppFiat struct {
	ent.Schema
}

func (AppFiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppFiat.
func (AppFiat) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			JSON("display_names", []string{}).
			Optional().
			Default([]string{}),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			Bool("disabled").
			Optional().
			Default(false),
		field.
			Bool("display").
			Optional().
			Default(true),
		field.
			Uint32("display_index").
			Optional().
			Default(0),
	}
}

// Edges of the AppFiat.
func (AppFiat) Edges() []ent.Edge {
	return nil
}
