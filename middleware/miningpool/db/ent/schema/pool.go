package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
)

// Pool holds the schema definition for the Pool entity.
type Pool struct {
	ent.Schema
}

func (Pool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Pool.
func (Pool) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("mining_pool_type").Optional().Default(""),
		field.
			String("name").Optional().Default(""),
		field.
			String("site").Optional().Default(""),
		field.
			String("logo").Optional().Default(""),
		field.
			String("description").Optional().Default(""),
	}
}

// Edges of the Pool.
func (Pool) Edges() []ent.Edge {
	return nil
}
