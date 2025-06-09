package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// OrderUser holds the schema definition for the OrderUser entity.
type OrderUser struct {
	ent.Schema
}

func (OrderUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OrderUser.
func (OrderUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("name").Optional().Default(""),
		field.
			Text("read_page_link").Optional().Default(""),
	}
}

// Edges of the OrderUser.
func (OrderUser) Edges() []ent.Edge {
	return nil
}

func (OrderUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("good_user_id"),
	}
}
