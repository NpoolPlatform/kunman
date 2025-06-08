package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	types "github.com/NpoolPlatform/kunman/message/basetypes/chain/v1"
	"github.com/google/uuid"
)

// CoinUsedFor holds the schema definition for the CoinUsedFor entity.
type CoinUsedFor struct {
	ent.Schema
}

func (CoinUsedFor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CoinUsedFor.
func (CoinUsedFor) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("used_for").
			Optional().
			Default(types.CoinUsedFor_DefaultCoinUsedFor.String()),
		field.
			Uint32("priority").
			Optional().
			Default(1),
	}
}

// Edges of the CoinUsedFor.
func (CoinUsedFor) Edges() []ent.Edge {
	return nil
}
