package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// DelegatedStaking holds the schema definition for the DelegatedStaking entity.
type DelegatedStaking struct {
	ent.Schema
}

func (DelegatedStaking) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the DelegatedStaking.
func (DelegatedStaking) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("contract_code_url").
			Optional().
			Default(""),
		field.
			String("contract_code_branch").
			Optional().
			Default(""),
		field.
			String("contract_state").
			Optional().
			Default(types.ContractState_ContractWaitDeployment.String()),
	}
}

// Edges of the DelegatedStaking.
func (DelegatedStaking) Edges() []ent.Edge {
	return nil
}

func (DelegatedStaking) Indexes() []ent.Index {
	return nil
}
