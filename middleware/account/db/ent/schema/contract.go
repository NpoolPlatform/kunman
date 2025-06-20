package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// Contract holds the schema definition for the Contract entity.
type Contract struct {
	ent.Schema
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Contract.
func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("delegated_staking_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("account_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("backup").
			Optional().
			Default(false),
		field.
			String("contract_operator_type").
			Optional().
			Default(basetypes.ContractOperatorType_DefaultContractOperatorType.String()),
	}
}

// Edges of the Contract.
func (Contract) Edges() []ent.Edge {
	return nil
}
