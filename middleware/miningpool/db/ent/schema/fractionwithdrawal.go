package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// FractionWithdrawal holds the schema definition for the FractionWithdrawal entity.
type FractionWithdrawal struct {
	ent.Schema
}

func (FractionWithdrawal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FractionWithdrawal.
func (FractionWithdrawal) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("order_user_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).Optional().
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("fraction_withdraw_state").Optional().Default(""),
		field.
			Uint32("withdraw_at").Optional().Default(0),
		field.
			Uint32("promise_pay_at").Optional().Default(0),
		field.
			Text("msg").Optional().Default(""),
	}
}

// Edges of the FractionWithdrawal.
func (FractionWithdrawal) Edges() []ent.Edge {
	return nil
}

func (FractionWithdrawal) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id"),
	}
}
