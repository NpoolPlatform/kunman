package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

// CreditAllocated holds the schema definition for the CreditAllocated entity.
type CreditAllocated struct {
	ent.Schema
}

func (CreditAllocated) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CreditAllocated.
func (CreditAllocated) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("extra").
			Optional().
			Default("").
			MaxLen(512), //nolint
	}
}

// Edges of the CreditAllocated.
func (CreditAllocated) Edges() []ent.Edge {
	return nil
}
