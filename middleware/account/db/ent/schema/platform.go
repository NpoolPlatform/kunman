package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
)

// Platform holds the schema definition for the Platform entity.
type Platform struct {
	ent.Schema
}

func (Platform) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Platform.
func (Platform) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("used_for").
			Optional().
			Default(basetypes.AccountUsedFor_DefaultAccountUsedFor.String()),
		field.
			Bool("backup").
			Optional().
			Default(false),
	}
}

// Edges of the Platform.
func (Platform) Edges() []ent.Edge {
	return nil
}
