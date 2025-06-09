package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// RootUser holds the schema definition for the RootUser entity.
type RootUser struct {
	ent.Schema
}

func (RootUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the RootUser.
func (RootUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").Optional().Default(""),
		field.
			UUID("pool_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("email").Optional().Default(""),
		field.
			Text("auth_token").Optional().Default(""),
		field.
			Text("auth_token_salt").Optional().Default(""),
		field.
			Bool("authed").Optional().Default(false),
		field.
			String("remark").Optional().Default(""),
	}
}

// Edges of the RootUser.
func (RootUser) Edges() []ent.Edge {
	return nil
}
