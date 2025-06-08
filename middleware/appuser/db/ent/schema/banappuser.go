package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// BanAppUser holds the schema definition for the BanAppUser entity.
type BanAppUser struct {
	ent.Schema
}

func (BanAppUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the BanAppUser.
func (BanAppUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("message").
			Default(""),
	}
}

// Edges of the BanAppUser.
func (BanAppUser) Edges() []ent.Edge {
	return nil
}
