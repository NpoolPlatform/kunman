package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// TopMost holds the schema definition for the TopMost entity.
type TopMost struct {
	ent.Schema
}

func (TopMost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the TopMost.
func (TopMost) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("top_most_type").
			Optional().
			Default(types.GoodTopMostType_DefaultGoodTopMostType.String()),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			String("message").
			Optional().
			Default(""),
		field.
			String("target_url").
			Optional().
			Default(""),
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

// Edges of the TopMost.
func (TopMost) Edges() []ent.Edge {
	return nil
}
