package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// OrderLock holds the schema definition for the OrderLock entity.
type OrderLock struct {
	ent.Schema
}

func (OrderLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the OrderLock.
func (OrderLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		// For commission lock
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("lock_type").
			Optional().
			Default(types.OrderLockType_DefaultOrderLockType.String()),
	}
}

// Edges of the OrderLock.
func (OrderLock) Edges() []ent.Edge {
	return nil
}

func (OrderLock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
