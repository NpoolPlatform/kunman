package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// OutOfGas holds the schema definition for the OutOfGas entity.
type OutOfGas struct {
	ent.Schema
}

func (OutOfGas) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OutOfGas.
func (OutOfGas) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
	}
}

// Edges of the OutOfGas.
func (OutOfGas) Edges() []ent.Edge {
	return nil
}

func (OutOfGas) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
