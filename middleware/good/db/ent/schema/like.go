package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

func (Like) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Bool("like"),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return nil
}

func (Like) Indexes() []ent.Index {
	return nil
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_good_likes"},
	}
}
