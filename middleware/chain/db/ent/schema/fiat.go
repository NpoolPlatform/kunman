package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
)

// Fiat holds the schema definition for the Fiat entity.
type Fiat struct {
	ent.Schema
}

func (Fiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Fiat.
func (Fiat) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("unit").
			Optional().
			Default(""),
	}
}

// Edges of the Fiat.
func (Fiat) Edges() []ent.Edge {
	return nil
}
