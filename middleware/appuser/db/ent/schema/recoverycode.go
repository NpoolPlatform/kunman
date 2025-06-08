package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// RecoveryCode holds the schema definition for the RecoveryCode entity.
type RecoveryCode struct {
	ent.Schema
}

func (RecoveryCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the RecoveryCode.
func (RecoveryCode) Fields() []ent.Field {
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
			String("code").
			Optional().
			Default(""),
		field.
			Bool("used").
			Optional().
			Default(false),
	}
}

// Edges of the RecoveryCode.
func (RecoveryCode) Edges() []ent.Edge {
	return nil
}
