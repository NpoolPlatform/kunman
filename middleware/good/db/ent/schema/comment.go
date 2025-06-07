package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
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
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("content").
			Optional().
			Default(""),
		field.
			UUID("reply_to_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("anonymous").
			Optional().
			Default(false),
		field.
			Bool("trial_user").
			Optional().
			Default(false),
		field.
			Bool("purchased_user").
			Optional().
			Default(false),
		field.
			Bool("hide").
			Optional().
			Default(false),
		field.
			String("hide_reason").
			Optional().
			Default(types.GoodCommentHideReason_DefaultGoodCommentHideReason.String()),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_good_id"),
	}
}

func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_good_comments"},
	}
}
