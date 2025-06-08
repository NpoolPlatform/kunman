package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

// ChainBase holds the schema definition for the ChainBase entity.
type ChainBase struct {
	ent.Schema
}

func (ChainBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the ChainBase.
func (ChainBase) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("native_unit").
			Optional().
			Default(""),
		field.
			String("atomic_unit").
			Optional().
			Default(""),
		field.
			Uint32("unit_exp").
			Optional().
			Default(0),
		field.
			String("env").
			Optional().
			Default(""),
		field.
			String("chain_id").
			Optional().
			Default(""),
		field.
			String("nickname").
			Optional().
			Default(""),
		field.
			String("gas_type").
			Optional().
			Default(basetypes.GasType_DefaultGasType.String()),
	}
}

// Edges of the ChainBase.
func (ChainBase) Edges() []ent.Edge {
	return nil
}
