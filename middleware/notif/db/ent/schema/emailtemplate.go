package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"

	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

// EmailTemplate holds the schema definition for the EmailTemplate entity.
type EmailTemplate struct {
	ent.Schema
}

func (EmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the EmailTemplate.
func (EmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("lang_id", uuid.UUID{}),
		field.
			String("default_to_username"),
		field.
			String("used_for").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			String("sender").
			Optional().
			Default(""),
		field.
			JSON("reply_tos", []string{}).
			Optional().
			Default([]string{}),
		field.
			JSON("cc_tos", []string{}).
			Optional().
			Default([]string{}),
		field.
			String("subject").
			Optional().
			Default(""),
		field.
			Text("body").
			Optional().
			Default(""),
	}
}
