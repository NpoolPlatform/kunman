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

// UnsoldStatement holds the schema definition for the UnsoldStatement entity.
type UnsoldStatement struct {
	ent.Schema
}

func (UnsoldStatement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the UnsoldStatement.
func (UnsoldStatement) Fields() []ent.Field {
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
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			Uint32("benefit_date").
			Optional().
			Default(0),
		field.
			UUID("statement_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the UnsoldStatement.
func (UnsoldStatement) Edges() []ent.Edge {
	return nil
}

func (UnsoldStatement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mining_unsolds"},
	}
}
