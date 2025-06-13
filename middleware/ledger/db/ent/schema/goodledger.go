//nolint:nolintlint,dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodLedger holds the schema definition for the GoodLedger entity.
type GoodLedger struct {
	ent.Schema
}

func (GoodLedger) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the GoodLedger.
func (GoodLedger) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
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
		field.Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.Float("to_platform").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.Float("to_user").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
	}
}

// Edges of the GoodLedger.
func (GoodLedger) Edges() []ent.Edge {
	return nil
}

func (GoodLedger) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mining_generals"},
	}
}
