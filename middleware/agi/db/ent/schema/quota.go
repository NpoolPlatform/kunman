package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Quota holds the schema definition for the Quota entity.
type Quota struct {
	ent.Schema
}

func (Quota) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the Quota.
func (Quota) Fields() []ent.Field {
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
			Uint32("quota").
			Optional().
			Default(0),
		field.
			Uint32("consumed_quota").
			Optional().
			Default(0),
		field.
			Uint32("expired_at").
			Optional().
			Default(0),
	}
}

// Edges of the Quota.
func (Quota) Edges() []ent.Edge {
	return nil
}
