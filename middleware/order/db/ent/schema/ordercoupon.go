//nolint
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// OrderCoupon holds the schema definition for the OrderCoupon entity.
type OrderCoupon struct {
	ent.Schema
}

func (OrderCoupon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the OrderCoupon.
func (OrderCoupon) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("coupon_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
	}
}

// Edges of the OrderCoupon.
func (OrderCoupon) Edges() []ent.Edge {
	return nil
}

func (OrderCoupon) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
