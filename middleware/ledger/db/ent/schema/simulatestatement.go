package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SimulateStatement holds the schema definition for the SimulateStatement entity.
type SimulateStatement struct {
	ent.Schema
}

func (SimulateStatement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the SimulateStatement.
func (SimulateStatement) Fields() []ent.Field {
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
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("io_type").
			Optional().
			Default(basetypes.IOType_DefaultType.String()),
		field.
			String("io_sub_type").
			Optional().
			Default(basetypes.IOSubType_DefaultSubType.String()),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			String("io_extra").
			Optional().
			Default("").
			MaxLen(512), //nolint
		field.
			Bool("send_coupon").
			Optional().
			Default(false),
		field.
			Bool("cashable").
			Optional().
			Default(false),
	}
}

// Edges of the Detail.
func (SimulateStatement) Edges() []ent.Edge {
	return nil
}

func (SimulateStatement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "simulate_details"},
	}
}
