package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppUserThirdParty holds the schema definition for the AppUserThirdParty entity.
type AppUserThirdParty struct {
	ent.Schema
}

func (AppUserThirdParty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppUserThirdParty.
func (AppUserThirdParty) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("third_party_user_id").
			Optional().
			Default(""),
		field.
			UUID("third_party_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("third_party_username").
			Optional().
			Default(""),
		field.
			Text("third_party_avatar").
			Optional().
			Default(""),
	}
}

// Edges of the AppUserThirdParty.
func (AppUserThirdParty) Edges() []ent.Edge {
	return nil
}

func (AppUserThirdParty) Indexes() []ent.Index {
	return []ent.Index{}
}
