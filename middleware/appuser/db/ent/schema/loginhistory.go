package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

// LoginHistory holds the schema definition for the LoginHistory entity.
type LoginHistory struct {
	ent.Schema
}

func (LoginHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the LoginHistory.
func (LoginHistory) Fields() []ent.Field {
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
			String("client_ip").
			Optional().
			Default(""),
		field.
			String("user_agent").
			Optional().
			Default(""),
		field.
			String("location").
			Optional().
			Default(""),
		field.
			String("login_type").
			Optional().
			Default(basetypes.LoginType_FreshLogin.String()),
	}
}

// Edges of the LoginHistory.
func (LoginHistory) Edges() []ent.Edge {
	return nil
}
