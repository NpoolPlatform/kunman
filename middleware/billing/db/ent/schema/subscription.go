package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/mixin"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("package_name").
			Optional().
			Default(""),
		field.
			Other("usd_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("description").
			Optional().
			Default(""),
		field.
			Uint32("sort_order").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(0)
			}),
		field.
			String("package_type").
			Optional().
			Default(types.PackageType_DefaultPackageType.String()),
		field.
			Uint32("credit").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(0)
			}),
		field.
			String("reset_type").
			Optional().
			Default(types.ResetType_DefaultResetType.String()),
		field.
			Uint32("qps_limit").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(1)
			}),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return nil
}
