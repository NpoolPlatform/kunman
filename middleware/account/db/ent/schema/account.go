package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("address").
			Optional().
			Default(""),
		field.
			String("used_for").
			Optional().
			Default(basetypes.AccountUsedFor_DefaultAccountUsedFor.String()),
		field.
			Bool("platform_hold_private_key").
			Optional().
			Default(false),
		field.
			Bool("active").
			Optional().
			Default(true),
		field.
			Bool("locked").
			Optional().
			Default(false),
		field.
			String("locked_by").
			Optional().
			Default(basetypes.AccountLockedBy_DefaultLockedBy.String()),
		field.
			Bool("blocked").
			Optional().
			Default(false),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
