package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/mixin"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	"github.com/google/uuid"
)

// Exchange holds the schema definition for the Exchange entity.
type Exchange struct {
	ent.Schema
}

func (Exchange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Exchange.
func (Exchange) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("usage_type").
			Optional().
			Default(types.UsageType_DefaultUsageType.String()),
		field.
			Uint32("credit").
			Optional().
			Default(0),
		field.
			Uint32("exchange_threshold").
			Optional().
			Default(0),
		field.
			String("path").
			Optional().
			Default(""),
	}
}

// Edges of the Exchange.
func (Exchange) Edges() []ent.Edge {
	return nil
}
