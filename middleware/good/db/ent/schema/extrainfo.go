package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExtraInfo holds the schema definition for the ExtraInfo entity.
type ExtraInfo struct {
	ent.Schema
}

func (ExtraInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the ExtraInfo.
func (ExtraInfo) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("likes").
			Optional().
			Default(0),
		field.
			Uint32("dislikes").
			Optional().
			Default(0),
		field.
			Uint32("recommend_count").
			Optional().
			Default(0),
		field.
			Uint32("comment_count").
			Optional().
			Default(0),
		field.
			Uint32("score_count").
			Optional().
			Default(0),
		field.
			Other("score", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the ExtraInfo.
func (ExtraInfo) Edges() []ent.Edge {
	return nil
}

func (ExtraInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
