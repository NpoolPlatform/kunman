package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/mixin"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	"github.com/google/uuid"
)

// UserSubscription holds the schema definition for the UserSubscription entity.
type UserSubscription struct {
	ent.Schema
}

func (UserSubscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the UserSubscription.
func (UserSubscription) Fields() []ent.Field {
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
			UUID("package_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			String("usage_state").
			Optional().
			Default(types.UsageState_DefaultUsageState.String()),
		field.
			Uint32("subscription_credit").
			Optional().
			Default(0),
		field.
			Uint32("addon_credit").
			Optional().
			Default(0),
	}
}

// Edges of the UserSubscription.
func (UserSubscription) Edges() []ent.Edge {
	return nil
}
