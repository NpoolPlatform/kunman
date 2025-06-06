package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppSubscriptionOneShot holds the schema definition for the AppSubscriptionOneShot entity.
type AppSubscriptionOneShot struct {
	ent.Schema
}

func (AppSubscriptionOneShot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppSubscriptionOneShot.
func (AppSubscriptionOneShot) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("banner").
			Optional().
			Default(""),
		field.
			Other("usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the AppSubscriptionOneShot.
func (AppSubscriptionOneShot) Edges() []ent.Edge {
	return nil
}

func (AppSubscriptionOneShot) Indexes() []ent.Index {
	return nil
}
