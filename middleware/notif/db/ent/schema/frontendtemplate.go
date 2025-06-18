//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"

	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

// FrontendTemplate holds the schema definition for the FrontendTemplate entity.
type FrontendTemplate struct {
	ent.Schema
}

func (FrontendTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FrontendTemplate.
func (FrontendTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("lang_id", uuid.UUID{}),
		field.
			String("used_for").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Text("content").
			Optional().
			Default(""),
	}
}
