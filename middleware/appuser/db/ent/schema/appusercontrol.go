package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

// AppUserControl holds the schema definition for the AppUserControl entity.
type AppUserControl struct {
	ent.Schema
}

func (AppUserControl) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppUserControl.
func (AppUserControl) Fields() []ent.Field {
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
		field.Bool("signin_verify_by_google_authentication").
			Optional().
			Default(false),
		field.Bool("google_authentication_verified").
			Optional().
			Default(false),
		field.String("signin_verify_type").
			Optional().
			Default(basetypes.SignMethod_Email.String()),
		field.Bool("kol").
			Default(false),
		field.Bool("kol_confirmed").
			Default(false),
		field.
			UUID("selected_lang_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the AppUserControl.
func (AppUserControl) Edges() []ent.Edge {
	return nil
}

func (AppUserControl) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id").Unique(),
	}
}
