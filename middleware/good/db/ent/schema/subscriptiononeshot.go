package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SubscriptionOneShot holds the schema definition for the SubscriptionOneShot entity.
type SubscriptionOneShot struct {
	ent.Schema
}

func (SubscriptionOneShot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the SubscriptionOneShot.
func (SubscriptionOneShot) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("quota").
			Optional().
			Default(1080),
		field.
			Other("usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the SubscriptionOneShot.
func (SubscriptionOneShot) Edges() []ent.Edge {
	return nil
}

func (SubscriptionOneShot) Indexes() []ent.Index {
	return nil
}
