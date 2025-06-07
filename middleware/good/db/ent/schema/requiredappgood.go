//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// RequiredAppGood holds the schema definition for the RequiredAppGood entity.
type RequiredAppGood struct {
	ent.Schema
}

func (RequiredAppGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the RequiredAppGood.
func (RequiredAppGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("main_app_good_id", uuid.UUID{}),
		field.
			UUID("required_app_good_id", uuid.UUID{}),
		field.
			Bool("must").
			Optional().
			Default(false),
	}
}

// Edges of the RequiredAppGood.
func (RequiredAppGood) Edges() []ent.Edge {
	return nil
}

func (RequiredAppGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("main_app_good_id", "required_app_good_id"),
	}
}
