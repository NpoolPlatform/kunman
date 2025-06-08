package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppOAuthThirdParty holds the schema definition for the AppOAuthThirdParty entity.
type AppOAuthThirdParty struct {
	ent.Schema
}

func (AppOAuthThirdParty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppOAuthThirdParty.
func (AppOAuthThirdParty) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("third_party_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("client_id").
			Optional().
			Default(""),
		field.
			Text("client_secret").
			Optional().
			Default(""),
		field.
			String("callback_url").
			Optional().
			Default(""),
		field.
			String("salt").
			Optional().
			Default(""),
	}
}

// Edges of the AppOAuthThirdParty.
func (AppOAuthThirdParty) Edges() []ent.Edge {
	return nil
}
